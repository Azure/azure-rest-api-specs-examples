Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.3.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

```go
package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Create.json
func ExampleAFDOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		armcdn.AFDOriginGroup{
			Properties: &armcdn.AFDOriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Int32Ptr(10),
					ProbePath:              to.StringPtr("<probe-path>"),
					ProbeProtocol:          armcdn.ProbeProtocolNotSet.ToPtr(),
					ProbeRequestType:       armcdn.HealthProbeRequestTypeNotSet.ToPtr(),
				},
				LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
					AdditionalLatencyInMilliseconds: to.Int32Ptr(1000),
					SampleSize:                      to.Int32Ptr(3),
					SuccessfulSamplesRequired:       to.Int32Ptr(3),
				},
				TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Int32Ptr(5),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AFDOriginGroupsClientCreateResult)
}
```
