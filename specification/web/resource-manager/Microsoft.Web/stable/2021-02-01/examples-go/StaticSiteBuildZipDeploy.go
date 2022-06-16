package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/StaticSiteBuildZipDeploy.json
func ExampleStaticSitesClient_BeginCreateZipDeploymentForStaticSiteBuild() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateZipDeploymentForStaticSiteBuild(ctx,
		"<resource-group-name>",
		"<name>",
		"<environment-name>",
		armappservice.StaticSiteZipDeploymentARMResource{
			Properties: &armappservice.StaticSiteZipDeployment{
				APIZipURL:        to.StringPtr("<apizip-url>"),
				AppZipURL:        to.StringPtr("<app-zip-url>"),
				DeploymentTitle:  to.StringPtr("<deployment-title>"),
				FunctionLanguage: to.StringPtr("<function-language>"),
				Provider:         to.StringPtr("<provider>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
