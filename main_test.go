package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
)
func TestIntegration(t *testing.T) {
    ts := httptest.NewServer(http.HandlerFunc(Handler))
    defer ts.Close()

    res, err := http.Get(ts.URL)
    if err != nil { t.Fatal(err) }
    if res.StatusCode != http.StatusOK {
        t.Errorf("Expected 200, got %d", res.StatusCode)
    }
}