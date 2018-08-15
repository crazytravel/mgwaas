package mgw

import (
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

func loadTemplate() string {
	pathStr := "./resources/mgw.tpl"
	mgwPath, err := filepath.Abs(pathStr)
	if err != nil {
		log.Panicln(err.Error())
		return ""
	}
	mgwData, err := ioutil.ReadFile(mgwPath)
	if err != nil {
		log.Panicln(err.Error())
		return ""
	}
	return string(mgwData)
}

// ParseTemplate is to parse customize template
func ParseTemplate(wr io.Writer) {
	tmpl, err := template.New("mgw").Parse(loadTemplate())
	if err != nil {
		log.Panicln(err.Error())
		return
	}
	mgwCxt := createMockData()
	tmpl.Execute(wr, mgwCxt)
}

func createMockData() *MicroGatewayTemplate {
	return &MicroGatewayTemplate{
		EdgeConfig: &EdgeConfig{
			Bootstrap:        "https://edgemicroservices.apigee.net/edgemicro/bootstrap/organization/wangshuo866-eval/environment/test",
			JwtPublicKey:     "https://wangshuo866-eval-test.apigee.net/edgemicro-auth/publicKey",
			ManagementURI:    "https://api.enterprise.apigee.com",
			VaultName:        "microgateway",
			AuthURI:          "https://%s-%s.apigee.net/edgemicro-auth",
			BaseURI:          "https://edgemicroservices.apigee.net/edgemicro/%s/organization/%s/environment/%s",
			BootstrapMessage: "Please copy the following property to the edge micro agent config",
			KeySecretMessage: "The following credentials are required to start edge micro",
			Products:         "https://wangshuo866-eval-test.apigee.net/edgemicro-auth/products",
		},
		EdgeMicro: &EdgeMicro{
			Port:                     8000,
			MaxConnections:           1000,
			ConfigChangePollInterval: 600,
			Logging: &Logging{
				Level:            "error",
				Dir:              "/var/tmp",
				StatsLogInterval: 60,
				RotateInterval:   24,
			},
			Plugins: &Plugins{
				Sequence: []string{"oauth"},
			},
		},
		Headers: &Headers{
			XForwardedFor:  true,
			XForwardedHost: true,
			XRequestID:     true,
			XResponseTime:  true,
			Via:            true,
		},
		OAuth: &OAuth{
			AllowNoAuthorization:      false,
			AllowInvalidAuthorization: false,
			VerifyAPIKeyURL:           "https://wangshuo866-eval-test.apigee.net/edgemicro-auth/verifyApiKey",
		},
		Analytics: &Analytics{
			URI:           "https://edgemicroservices.apigee.net/edgemicro/axpublisher/organization/wangshuo866-eval/environment/test",
			BufferSize:    10000,
			BatchSize:     500,
			FlushInterval: 5000,
		},
	}
}
