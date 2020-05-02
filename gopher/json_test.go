package main

import "testing"

func TestToJson(t *testing.T) {
	ts := []struct {
		uuid string
		fib  uint64
	}{
		{uuid: "52fdfc07-2182-454f-963f-5f0f9a621d72", fib: 0},
		{uuid: "9566c74d-1003-4c4d-bbbb-0407d1e2c649", fib: 1},
		{uuid: "81855ad8-681d-4d86-91e9-1e00167939cb", fib: 2},
		{uuid: "6694d2c4-22ac-4208-a007-2939487f6999", fib: 3},
		{uuid: "eb9d18a4-4784-445d-87f3-c67cf22746e9", fib: 21},
		{uuid: "95af5a25-3679-41ba-a2ff-6cd471c483f1", fib: 987},
		{uuid: "680b4e7c-8b76-4a1b-9d49-d4955c848621", fib: 6765},
	}
	tr := []string{
		`{"uuid":"52fdfc07-2182-454f-963f-5f0f9a621d72","fib":"0"}`,
		`{"uuid":"9566c74d-1003-4c4d-bbbb-0407d1e2c649","fib":"1"}`,
		`{"uuid":"81855ad8-681d-4d86-91e9-1e00167939cb","fib":"2"}`,
		`{"uuid":"6694d2c4-22ac-4208-a007-2939487f6999","fib":"3"}`,
		`{"uuid":"eb9d18a4-4784-445d-87f3-c67cf22746e9","fib":"21"}`,
		`{"uuid":"95af5a25-3679-41ba-a2ff-6cd471c483f1","fib":"987"}`,
		`{"uuid":"680b4e7c-8b76-4a1b-9d49-d4955c848621","fib":"6765"}`,
	}
	for i, tc := range ts {
		r := string(toJson([]byte(tc.uuid), tc.fib))
		if tr[i] != r {
			t.Errorf("scenario %d:\nexpected:\t%s\ngot:\t\t%s", i, tr[i], r)
		}
	}
}

var benchUuid = []byte("9566c74d-1003-4c4d-bbbb-0407d1e2c649")

func BenchmarkToJsonAt2(b *testing.B) {
	benchmarkToJsonAt(b, 2)
}

func BenchmarkToJsonAt4(b *testing.B) {
	benchmarkToJsonAt(b, 4)
}

func BenchmarkToJsonAt8(b *testing.B) {
	benchmarkToJsonAt(b, 8)
}

func BenchmarkToJsonAt16(b *testing.B) {
	benchmarkToJsonAt(b, 16)
}

func benchmarkToJsonAt(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		toJson(benchUuid, n)
	}
}
