package e2e

import (
	"fmt"
	"log"
	"testing"

	"github.com/sclevine/agouti"
)

var (
	webDriver *agouti.WebDriver
)

func TestMain(m *testing.M) {
	if err := testMain(m); err != nil {
		log.Fatalf("err = %v\n", err)
	}
}

func testMain(m *testing.M) error {
	// Setup WebDriver
	webDriver = agouti.ChromeDriver()
	if err := webDriver.Start(); err != nil {
		return err
	}
	defer webDriver.Stop()

	// Run tests
	if status := m.Run(); status != 0 {
		return fmt.Errorf("Run status = %v", status)
	}

	return nil
}

func TestIndex(t *testing.T) {
	page, err := webDriver.NewPage()
	if err != nil {
		t.Error(err)
	}
	if err := page.Navigate("https://github.com/"); err != nil {
		t.Error(err)
	}
	text, _ := page.Find(".alt-h0.text-white.lh-condensed-ultra.mb-3").Text()
	if expected, actual := "Built for developers", text; expected != actual {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}
