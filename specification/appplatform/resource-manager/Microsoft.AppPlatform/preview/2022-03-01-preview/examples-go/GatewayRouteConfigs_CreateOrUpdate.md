```go
package armappplatform_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-03-01-preview/examples/GatewayRouteConfigs_CreateOrUpdate.json
func ExampleGatewayRouteConfigsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappplatform.NewGatewayRouteConfigsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<gateway-name>",
		"<route-config-name>",
		armappplatform.GatewayRouteConfigResource{
			Properties: &armappplatform.GatewayRouteConfigProperties{
				AppResourceID: to.Ptr("<app-resource-id>"),
				Routes: []*armappplatform.GatewayAPIRoute{
					{
						Filters: []*string{
							to.Ptr("StripPrefix=2"),
							to.Ptr("RateLimit=1,1s")},
						Predicates: []*string{
							to.Ptr("Path=/api5/customer/**")},
						SsoEnabled: to.Ptr(true),
						Title:      to.Ptr("<title>"),
					}},
			},
		},
		&armappplatform.GatewayRouteConfigsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv0.5.0/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.
