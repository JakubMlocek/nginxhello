package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"reflect"
	"testing"
	"time"
)

func TestTerraformHellonginxExample(t *testing.T) {
	terrafromOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/nginxhello",
		Vars: map[string]interface{}{
			"docker_image": "jakmlo/nginxhello",
			"port":         "8081",
		},
	})

	terraform.InitAndApply(t, terrafromOptions)

	defer terraform.Destroy(t, terrafromOptions)

	port := terraform.Output(t, terrafromOptions, "port")
	fmt.Println(reflect.TypeOf(port))

	url := fmt.Sprintf("http://localhost:%s", port)
	properSource :=
		`<html>
    <head>
        <title>Hello Word!</title>
    </head>
    <body>
        <h1>I am  working!</h1>
    </body>
</html>`

	fmt.Println(properSource)
	http_helper.HttpGetWithRetry(t, url, nil, 200, properSource, 10, 5*time.Second)
}
