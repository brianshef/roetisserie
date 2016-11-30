package config

import (
    "fmt"
    "os"
)

const (
    fbClientIdEnvVar        string  = "ROETISSERIE_FB_CLIENT_ID"
    fbClientSecretEnvVar    string  = "ROETISSERIE_FB_CLIENT_SECRET"
)

var Config Configuration

type facebookAppConfiguration struct {
    ClientId        string
    ClientSecret    string
    // redirectUrl     string
    // scopes          []string
    // endpoint        string
}

type Configuration struct {
    FacebookApp facebookAppConfiguration
}

func LoadFacebookAppConfig() Configuration {
    clientId    := os.Getenv(fbClientIdEnvVar)
    clientSecret := os.Getenv(fbClientSecretEnvVar)

    fmt.Println("Read Facebook ClientId: ", clientId )
    fmt.Println("Read Facebook ClientSecret: ", clientSecret )

    Config := Configuration{
        FacebookApp: facebookAppConfiguration{
            ClientId: clientId,
            ClientSecret: clientSecret,
        },
    }

    fmt.Println("Loaded Facebook ClientId: ", Config.FacebookApp.ClientId )
    fmt.Println("Loaded Facebook ClientSecret: ", Config.FacebookApp.ClientSecret )

    return Config
}
