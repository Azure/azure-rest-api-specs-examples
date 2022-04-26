Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftrafficmanager%2Farmtrafficmanager%2Fv0.4.0/sdk/resourcemanager/trafficmanager/armtrafficmanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PUT-MultiValue.json
func ExampleProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewProfilesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		armtrafficmanager.Profile{
			Location: to.Ptr("<location>"),
			Properties: &armtrafficmanager.ProfileProperties{
				DNSConfig: &armtrafficmanager.DNSConfig{
					RelativeName: to.Ptr("<relative-name>"),
					TTL:          to.Ptr[int64](35),
				},
				MaxReturn: to.Ptr[int64](2),
				MonitorConfig: &armtrafficmanager.MonitorConfig{
					Path:     to.Ptr("<path>"),
					Port:     to.Ptr[int64](80),
					Protocol: to.Ptr(armtrafficmanager.MonitorProtocolHTTP),
				},
				ProfileStatus:               to.Ptr(armtrafficmanager.ProfileStatusEnabled),
				TrafficRoutingMethod:        to.Ptr(armtrafficmanager.TrafficRoutingMethodMultiValue),
				TrafficViewEnrollmentStatus: to.Ptr(armtrafficmanager.TrafficViewEnrollmentStatusDisabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
