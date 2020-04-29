package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/angelospillos/microwars/rusty_nail/lib"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	statusOK     = "{\"status\": \"ok\"}"
	statusError  = "{\"status\": \"not connected\"}"
	resultStart  = "{ \"uuid\": \""
	resultMiddle = "\", \"fib\": "
	resultEnd    = "}"
)

var (
	punchesTaken = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "taken",
		Help: "I need a hero",
	}, []string{"punchType"})

	punchesGiven = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "given",
		Help: "I need a zero",
	}, []string{"punchType"})

	knockouts = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "knockouts",
		Help: "IMy name is nero",
	}, []string{"punchType"})

	opponentURL         = "http://localhost:8081"
	refURL              = "http://localhost:8082"
	opponentStatusURL   = fmt.Sprintf("%s/%s", opponentURL, "status")
	opponentJabURL      = fmt.Sprintf("%s/%s", opponentURL, "jab")
	opponentCrossURL    = fmt.Sprintf("%s/%s", opponentURL, "cross")
	opponentHookURL     = fmt.Sprintf("%s/%s", opponentURL, "hook")
	opponentUppercutURL = fmt.Sprintf("%s/%s", opponentURL, "uppercut")

	client = http.Client{
		Timeout: 4 * time.Second,
	}
)

func main() {
	if os.Getenv("OPPONENT_URL") != "" {
		opponentURL = os.Getenv("OPPONENT_URL")
	}

	if os.Getenv("REF_URL") != "" {
		refURL = os.Getenv("REF_URL")
	}

	opponentStatusURL = fmt.Sprintf("%s/%s", opponentURL, "status")
	opponentJabURL = fmt.Sprintf("%s/%s", opponentURL, "jab")
	opponentCrossURL = fmt.Sprintf("%s/%s", opponentURL, "cross")
	opponentHookURL = fmt.Sprintf("%s/%s", opponentURL, "hook")
	opponentUppercutURL = fmt.Sprintf("%s/%s", opponentURL, "uppercut")

	fmt.Printf("Using OPPONENT_URL=%s\n", opponentURL)
	fmt.Printf("Using REF_URL=%s\n", refURL)

	recoverConfig := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}

	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
		Prefork:               true,
	})
	app.Use(recover.New(recoverConfig))

	app.Get("/", home)
	app.Get("/status", status)
	app.Get("/test", test)
	app.Get("/work", work)
	app.Get("/combat", combat)
	app.Get("/jab", jab)
	app.Get("/hook", hook)
	app.Get("/cross", cross)
	app.Get("/uppercut", uppercut)
	app.Listen(8080)
}

func getDaRef(url string, path string) {
	fmt.Printf("Calling ref:%s%s\n", url, path)
	var req strings.Builder
	req.WriteString(url)
	req.WriteString(path)
	resp, err := client.Get(req.String())
	if err != nil {
		fmt.Printf("Error calling ref: %+v\n", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Printf("Error calling ref: status=%d\n", resp.StatusCode)
	}
}

func getYourRoarOn(url string, callref bool) error {
	fmt.Print("Hit me baby, ", url, "\n")
	punchesTaken.With(prometheus.Labels{
		"punchType": url,
	}).Inc()

	resp, err := client.Get(url)
	if err != nil {
		if callref {
			go getDaRef(refURL, "/won?err=1")
		}

		knockouts.With(prometheus.Labels{
			"punchType": url,
		}).Inc()

		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 && callref {
		go getDaRef(refURL, fmt.Sprintf("/won?got-status=%d", resp.StatusCode))
		knockouts.With(prometheus.Labels{
			"punchType": url,
		}).Inc()
		return errors.New("Got status " + resp.Status)
	}

	return nil
}

func doDaRoar(times uint16) string {
	uuid := uuid.Must(uuid.NewRandom()).String()
	var req strings.Builder
	req.WriteString(resultStart)
	req.WriteString(uuid)
	req.WriteString(resultMiddle)
	req.WriteString(fmt.Sprintf("%d", lib.Fib(times)))
	req.WriteString(resultEnd)
	return req.String()
}

func home(c *fiber.Ctx) {
	c.Send("Reporting for duty")
}

func status(c *fiber.Ctx) {
	c.Send(statusOK)
}

func test(c *fiber.Ctx) {
	err := getYourRoarOn(opponentStatusURL, false)
	if err != nil {
		fmt.Printf("Error testing: %+v\n", err)
		c.Send(statusError)
		return
	}
	c.Send(statusOK)
}

func work(c *fiber.Ctx) {
	c.Send(doDaRoar(20))
}

func combat(c *fiber.Ctx) {
	go getYourRoarOn(opponentJabURL, true)
	go getYourRoarOn(opponentHookURL, true)
	c.Send(statusOK)
}

func jab(c *fiber.Ctx) {
	punchesTaken.With(prometheus.Labels{
		"punchType": c.Path(),
	}).Inc()
	go getYourRoarOn(opponentJabURL, true)
	go getYourRoarOn(opponentJabURL, true)
	c.Send(doDaRoar(2))
}

func cross(c *fiber.Ctx) {
	punchesTaken.With(prometheus.Labels{
		"punchType": c.Path(),
	}).Inc()
	go getYourRoarOn(opponentJabURL, true)
	go getYourRoarOn(opponentJabURL, true)
	go getYourRoarOn(opponentCrossURL, true)
	c.Send(doDaRoar(4))

}

func hook(c *fiber.Ctx) {
	punchesTaken.With(prometheus.Labels{
		"punchType": c.Path(),
	}).Inc()
	go getYourRoarOn(opponentHookURL, true)
	go getYourRoarOn(opponentHookURL, true)
	go getYourRoarOn(opponentUppercutURL, true)
	c.Send(doDaRoar(8))
}

func uppercut(c *fiber.Ctx) {
	punchesTaken.With(prometheus.Labels{
		"punchType": c.Path(),
	}).Inc()
	go getYourRoarOn(opponentCrossURL, true)
	go getYourRoarOn(opponentHookURL, true)
	go getYourRoarOn(opponentUppercutURL, true)
	c.Send(doDaRoar(16))
}
