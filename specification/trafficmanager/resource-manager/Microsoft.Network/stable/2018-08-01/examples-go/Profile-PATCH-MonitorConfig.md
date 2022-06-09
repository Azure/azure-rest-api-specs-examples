```go
package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PATCH-MonitorConfig.json
func ExampleProfilesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"azuresdkfornetautoresttrafficmanager2583",
		"azuresdkfornetautoresttrafficmanager6192",
		armtrafficmanager.Profile{
			Properties: &armtrafficmanager.ProfileProperties{
				MonitorConfig: &armtrafficmanager.MonitorConfig{
					Path: to.Ptr("/testpath.aspx"),
					CustomHeaders: []*armtrafficmanager.MonitorConfigCustomHeadersItem{
						{
							Name:  to.Ptr("header-1"),
							Value: to.Ptr("value-1"),
						},
						{
							Name:  to.Ptr("header-2"),
							Value: to.Ptr("value-2"),
						}},
					IntervalInSeconds:         to.Ptr[int64](30),
					Port:                      to.Ptr[int64](80),
					TimeoutInSeconds:          to.Ptr[int64](6),
					ToleratedNumberOfFailures: to.Ptr[int64](4),
					Protocol:                  to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftrafficmanager%2Farmtrafficmanager%2Fv1.0.0/sdk/resourcemanager/trafficmanager/armtrafficmanager/README.md) on how to add the SDK to your project and authenticate.
