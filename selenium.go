package mangaworld

import (
	"errors"
	"io"
	"os"
	"github.com/tebeka/selenium"
)

type ServiceConfig struct {
	SeleniumPath    string
	GeckoDriverPath string
	Port            int
	Verbose         bool
	Output          io.Writer
}

type WebDriverConfig struct {
	Capabilities selenium.Capabilities
	UrlService   string
}

var service *selenium.Service
var serviceActive bool

func NewService(cfg *ServiceConfig) (err error) {

	if serviceActive {
		return errors.New("Service already exist")
	}

	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),               // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(cfg.GeckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	}

	if cfg.Output != nil {
		opts = append(opts, selenium.Output(os.Stderr))
	}

	selenium.SetDebug(cfg.Verbose)

	service, err = selenium.NewSeleniumService(cfg.SeleniumPath, cfg.Port, opts...)
	if err != nil {
		return err
	}

	serviceActive = true

	return nil
}

func NewDefaultService() error {

	cfg := ServiceConfig{
		SeleniumPath:    "/home/kiritonya/go/src/github.com/tebeka/selenium/vendor/selenium-server.jar",
		GeckoDriverPath: "/home/kiritonya/go/src/github.com/tebeka/selenium/vendor/geckodriver",
		Port:            4444,
		Verbose:         false,
		Output:          nil,
	}

	err := NewService(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func NewWebDriver(url string, cfg *WebDriverConfig) (selenium.WebDriver, error) {

	if !serviceActive {
		return nil, errors.New("Service is inactive, read the documentation.")
	}

	wd, err := selenium.NewRemote(cfg.Capabilities, cfg.UrlService)
	if err != nil {
		return nil, err
	}

	if err = wd.Get(url); err != nil {
		return nil, err
	}

	return wd, err
}

func NewDefaultWebDriver(url string) (selenium.WebDriver, error) {

	cfg := WebDriverConfig{
		Capabilities: selenium.Capabilities{"browserName": "firefox"},
		UrlService:   "http://localhost:4444/wd/hub",
	}

	wd, err := NewWebDriver(url, &cfg)

	return wd, err
}

func ServiceIsActive() bool {
	return serviceActive
}

func doRequest(url string) (resp string, err error) {

	if !serviceActive {
		err = NewDefaultService()
		if err != nil {
			return "", err
		}
	}

	wd, err := NewDefaultWebDriver(url)
	if err != nil {
		return "", err
	}
	defer wd.Close()

	err = wd.Get(url)
	if err != nil {
		return "", err
	}

	pageString, err := wd.PageSource()
	if err != nil {
		return "", err
	}

	return pageString, nil
}
