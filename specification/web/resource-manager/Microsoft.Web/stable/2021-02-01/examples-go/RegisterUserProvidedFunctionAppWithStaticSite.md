Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.1.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/RegisterUserProvidedFunctionAppWithStaticSite.json
func ExampleStaticSitesClient_BeginRegisterUserProvidedFunctionAppWithStaticSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewStaticSitesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRegisterUserProvidedFunctionAppWithStaticSite(ctx,
		"<resource-group-name>",
		"<name>",
		"<function-app-name>",
		armappservice.StaticSiteUserProvidedFunctionAppARMResource{
			Properties: &armappservice.StaticSiteUserProvidedFunctionAppARMResourceProperties{
				FunctionAppRegion:     to.StringPtr("<function-app-region>"),
				FunctionAppResourceID: to.StringPtr("<function-app-resource-id>"),
			},
		},
		&armappservice.StaticSitesBeginRegisterUserProvidedFunctionAppWithStaticSiteOptions{IsForced: to.BoolPtr(true)})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("StaticSiteUserProvidedFunctionAppARMResource.ID: %s\n", *res.ID)
}
```
