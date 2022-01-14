Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftrafficmanager%2Farmtrafficmanager%2Fv0.2.0/sdk/resourcemanager/trafficmanager/armtrafficmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PATCH-MonitorConfig.json
func ExampleProfilesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtrafficmanager.NewProfilesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<profile-name>",
		armtrafficmanager.Profile{
			Properties: &armtrafficmanager.ProfileProperties{
				MonitorConfig: &armtrafficmanager.MonitorConfig{
					Path: to.StringPtr("<path>"),
					CustomHeaders: []*armtrafficmanager.MonitorConfigCustomHeadersItem{
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						},
						{
							Name:  to.StringPtr("<name>"),
							Value: to.StringPtr("<value>"),
						}},
					IntervalInSeconds:         to.Int64Ptr(30),
					Port:                      to.Int64Ptr(80),
					TimeoutInSeconds:          to.Int64Ptr(6),
					ToleratedNumberOfFailures: to.Int64Ptr(4),
					Protocol:                  armtrafficmanager.MonitorProtocol("HTTP").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProfilesClientUpdateResult)
}
```
