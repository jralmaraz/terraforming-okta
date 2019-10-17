package apps

import (
	//	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	//articulate "github.com/articulate/terraform-provider-okta/sdk"
	//	jsonParser "github.com/hashicorp/hcl/json"
	//"github.com/urfave/cli"
)

//func getApplications(d *schema.ResourceData, m interface{}) (*articulateOkta.AddressObj, error) {
func getApplications() {

	// Would use Okta's client and the ListApplications, but it is failing at the moment - issues linked below

	//ctx := context.Background()
	//client, _ := okta.NewClient(ctx, okta.WithOrgUrl(""), okta.WithToken(""))

	//https://github.com/okta/okta-sdk-golang/issues/97 - ListApplications will always fail to unmarshal from JSON
	//applications, resp, err := client.Application.ListApplications(nil)

	//	if resp == nil {
	//		log.Fatal(resp)
	//	}
	client := http.Client{Timeout: time.Second * 2}

	var OktaAPIToken = os.Getenv("OKTA_API_TOKEN")
	var OktaOrgName = os.Getenv("OKTA_ORG_NAME")
	var OktaBaseURL = os.Getenv("OKTA_BASE_URL")

	var url = "https://" + OktaOrgName + OktaBaseURL + "/api/v1/apps"

	request, reqerr := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "SSWS "+OktaAPIToken)

	if reqerr != nil {
		log.Fatal(reqerr)
	}

	resp, resperr := client.Do(request)

	if resperr != nil {
		log.Fatal(resperr)
	}

	if resperr != nil {
		log.Fatal(resperr)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	ioerr := ioutil.WriteFile("applications.json", body, os.ModePerm)
	if ioerr != nil {
		log.Fatal(ioerr)
	}

}
