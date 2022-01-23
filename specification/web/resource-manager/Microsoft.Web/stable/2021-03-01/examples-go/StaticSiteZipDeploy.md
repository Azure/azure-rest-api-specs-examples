Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.2.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StaticSiteZipDeploy.json
func ExampleStaticSitesClient_BeginCreateZipDeploymentForStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateZipDeploymentForStaticSite(ctx,
		"<resource-group-name>",
		"<name>",
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
```
