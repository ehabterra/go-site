package web_api

import (
	"net/http"
	"testing"

	"github.com/ehabterra/go-site/internal/environment"

	"github.com/ehabterra/go-site/internal/database"

	"github.com/gavv/httpexpect/v2"
)

func fastHTTPTester(t *testing.T) *httpexpect.Expect {
	db := database.Connect()
	server := NewServer(db)
	app := server.configure()
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(app.Handler()),
			Jar:       httpexpect.NewJar(),
		},
		// Report errors using testify.
		Reporter: httpexpect.NewAssertReporter(t),
	})
}

func TestFastHTTP(t *testing.T) {
	environment.LoadEnv("../../.env")
	e := fastHTTPTester(t)

	site := map[string]interface{}{
		"name": "span",
		"attributes": []map[string]interface{}{
			{
				"key":       "name",
				"value_str": "year",
				"value_int": nil,
			},
			{
				"key":       "value",
				"value_str": nil,
				"value_int": 2021,
			},
		},
	}

	siteID := e.POST("/site").WithJSON(site).Expect().
		Status(http.StatusOK).JSON().Object().Value("site_id").String().Raw()

	t.Log("new site_id: ", siteID)

	e.GET("/site/{id}", siteID).Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("site_id").
		ValueNotEqual("site_id", nil).
		ValueEqual("name", "span").
		Value("attributes").Array().Length().Equal(2)

	e.DELETE("/site/{id}", siteID).Expect().
		Status(http.StatusOK)

	e.GET("/site/{id}", siteID).Expect().
		Status(http.StatusBadRequest).
		JSON().Object().ValueEqual("success", false)
}
