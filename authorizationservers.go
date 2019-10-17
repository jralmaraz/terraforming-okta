package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	// okta "github.com/articulate/terraform-provider-okta/vendor/github.com/okta/okta-sdk-golang/okta" - use of vendored package not allowed
	articulate "github.com/articulate/terraform-provider-okta/sdk"
)

func main() {

	var OktaAPIToken = os.Getenv("OKTA_API_TOKEN")

	var OktaOrgName = os.Getenv("OKTA_ORG_NAME")

	var OktaBaseURL = os.Getenv("OKTA_BASE_URL")

	var OktaOrgURL = "https://" + OktaOrgName + OktaBaseURL

	//ctx := context.Background()

	//oktaclient, _ := okta.NewClient(ctx, okta.WithOrgUrl(OktaOrgURL), okta.WithToken(OktaAPIToken))

	httpclient := http.Client{Timeout: time.Second * 2}

	//I am struggling to correctly instantiate the APISupplement :(
	//oktaRequestExecutor := okta.NewRequestExecutor(&httpclient, nil, oktaclient.GetConfig())

	m := new(articulate.ApiSupplement)
	m.Token = OktaAPIToken
	m.BaseURL = OktaOrgURL
	m.Client = &httpclient
	//	fmt.Println(oktaclient.GetRequestExecutor())
	//m.RequestExecutor = oktaRequestExecutor

	var authzserver []map[string]interface{}

	authzservers, resp, err := m.ListAuthorizationServers()

	if authzservers == nil {
		fmt.Println(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &authzserver)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(authzserver)

	authzserversJSON, marshalerror := json.Marshal(body)

	if marshalerror != nil {
		log.Fatal(marshalerror)
	}

	fmt.Println(authzserversJSON)

}
