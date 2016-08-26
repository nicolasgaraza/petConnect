package tests

import "github.com/revel/revel/testing"

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestThatServicesGetWorks() {
	t.Get("/api/services/1")
	t.AssertOk()
	t.AssertContains("ServiceTest")
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
