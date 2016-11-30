// TODO - Build off of http://stackoverflow.com/questions/27368973/golang-facebook-authentication-using-golang-org-x-oauth2

package facebook

import (
    "fmt"
    fb "github.com/huandu/facebook"
    config "github.com/brianshef/roetisserie/config"
)

//  Facebook App Variables
var (
    FB_TEST_APP_ID                  = "169186383097898"
    FB_TEST_APP_SECRET              = "b2e4262c306caa3c7f5215d2d099b319"
    FB_TEST_VALID_ACCESS_TOKEN      = ""
    FB_TEST_VALID_SIGNED_REQUEST    = ""

    globalApp = fb.New(FB_TEST_APP_ID, FB_TEST_APP_SECRET)
)

var configuration config.Configuration

func configureFacebookApp() {
    configuration := config.LoadFacebookAppConfig();
    FB_TEST_APP_ID = configuration.FacebookApp.ClientId
    FB_TEST_APP_SECRET = configuration.FacebookApp.ClientSecret
}

func getAppAccessToken() {
    res, _ := fb.Get("/oauth/access_token", fb.Params{
        "client_id": FB_TEST_APP_ID,
        "client_secret": FB_TEST_APP_SECRET,
        "grant_type": "client_credentials",
    })
    fmt.Println("Got access token:", res)
}


func Test() {
    configureFacebookApp()
    getAppAccessToken()

    res, _ := fb.Get("/me", fb.Params{
        "fields": "first_name",
        "access_token": FB_TEST_VALID_ACCESS_TOKEN,
    })
    // fmt.Println("here is my facebook first name:", res["first_name"])
    fmt.Println(res)
}
