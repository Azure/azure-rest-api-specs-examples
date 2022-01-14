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

// x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Profile-PUT-MultiValue.json
func ExampleProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtrafficmanager.NewProfilesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		armtrafficmanager.Profile{
			Location: to.StringPtr("<location>"),
			Properties: &armtrafficmanager.ProfileProperties{
				DNSConfig: &armtrafficmanager.DNSConfig{
					RelativeName: to.StringPtr("<relative-name>"),
					TTL:          to.Int64Ptr(35),
				},
				MaxReturn: to.Int64Ptr(2),
				MonitorConfig: &armtrafficmanager.MonitorConfig{
					Path:     to.StringPtr("<path>"),
					Port:     to.Int64Ptr(80),
					Protocol: armtrafficmanager.MonitorProtocol("HTTP").ToPtr(),
				},
				ProfileStatus:               armtrafficmanager.ProfileStatus("Enabled").ToPtr(),
				TrafficRoutingMethod:        armtrafficmanager.TrafficRoutingMethod("MultiValue").ToPtr(),
				TrafficViewEnrollmentStatus: armtrafficmanager.TrafficViewEnrollmentStatus("Disabled").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProfilesClientCreateOrUpdateResult)
}
```
