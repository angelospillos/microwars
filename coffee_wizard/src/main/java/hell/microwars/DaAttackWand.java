package hell.microwars;

import com.wizzardo.http.framework.di.Service;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;
import java.util.concurrent.TimeUnit;

public class DaAttackWand implements Service {
    private static final String OPPONENT_URL;
    private static final String REFEREE_URL;

    static {
        if (System.getenv("OPPONENT_URL") != null) {
            OPPONENT_URL = System.getenv("OPPONENT_URL");
        } else {
            OPPONENT_URL = "https://enqfc8y2t9fo.x.pipedream.net";
        }
        if (System.getenv("REFEREE_URL") != null) {
            REFEREE_URL = System.getenv("REFEREE_URL");
        } else {
            REFEREE_URL = "https://enqfc8y2t9fo.x.pipedream.net";
        }
    }

    private final HttpClient opponentClient = HttpClient.newBuilder()
            .connectTimeout(Duration.ofSeconds(4))
            .build();

    private final HttpRequest opponentStatus = HttpRequest.newBuilder()
            .uri(new URI(OPPONENT_URL + "/status"))
            .GET()
            .build();
    private final HttpRequest opponentAttackJab = HttpRequest.newBuilder()
            .uri(new URI(OPPONENT_URL + "/jab"))
            .GET()
            .build();
    private final HttpRequest opponentAttackCross = HttpRequest.newBuilder()
            .uri(new URI(OPPONENT_URL + "/cross"))
            .GET()
            .build();
    private final HttpRequest opponentAttackHook = HttpRequest.newBuilder()
            .uri(new URI(OPPONENT_URL + "/hook"))
            .GET()
            .build();
    private final HttpRequest opponentAttackUppercut = HttpRequest.newBuilder()
            .uri(new URI(OPPONENT_URL + "/uppercut"))
            .GET()
            .build();

    private final HttpClient refereeClient = HttpClient.newBuilder().build();
    private final HttpRequest refereeStatusRequest = HttpRequest.newBuilder()
            .uri(new URI(REFEREE_URL))
            .GET()
            .build();
    private final HttpRequest refereeWon = HttpRequest.newBuilder()
            .uri(new URI(REFEREE_URL + "/won"))
            .POST(HttpRequest.BodyPublishers.ofString("{ \"name\": \"Coffee Wizard\" }"))
            .build();

    public DaAttackWand() throws URISyntaxException {
    }

    private void attack(final HttpRequest req) {
        opponentClient.sendAsync(req, HttpResponse.BodyHandlers.ofString())
                .completeOnTimeout(null, 4, TimeUnit.SECONDS)
                .thenAcceptAsync(x -> {
                    if (x == null) {
                        System.out.println("TIMEOUT! WON!");
                        try {
                            notifyRefereeWon();
                        } catch (IOException | InterruptedException e) {
                            e.printStackTrace();
                        }
                    }
                });
    }

    private void notifyRefereeWon() throws IOException, InterruptedException {
        refereeClient.send(refereeWon, HttpResponse.BodyHandlers.ofString());
    }


    public void jab() {
        attack(opponentAttackJab);
    }

    public void cross() {
        attack(opponentAttackCross);
    }

    public void hook() {
        attack(opponentAttackHook);
    }

    public void uppercut() {
        attack(opponentAttackUppercut);
    }

    public boolean connectionOk() throws IOException, InterruptedException {
        if (opponentClient.send(opponentStatus, HttpResponse.BodyHandlers.ofString()).statusCode() != 200) return false;
        return refereeClient.send(refereeStatusRequest, HttpResponse.BodyHandlers.ofString()).statusCode() == 200;
    }

}
