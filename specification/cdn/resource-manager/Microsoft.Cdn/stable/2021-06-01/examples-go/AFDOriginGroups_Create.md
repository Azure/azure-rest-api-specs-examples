Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcdn%2Farmcdn%2Fv0.5.0/sdk/resourcemanager/cdn/armcdn/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Create.json
func ExampleAFDOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewAFDOriginGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<origin-group-name>",
		armcdn.AFDOriginGroup{
			Properties: &armcdn.AFDOriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Ptr[int32](10),
					ProbePath:              to.Ptr("<probe-path>"),
					ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolNotSet),
					ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeNotSet),
				},
				LoadBalancingSettings: &armcdn.LoadBalancingSettingsParameters{
					AdditionalLatencyInMilliseconds: to.Ptr[int32](1000),
					SampleSize:                      to.Ptr[int32](3),
					SuccessfulSamplesRequired:       to.Ptr[int32](3),
				},
				TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: to.Ptr[int32](5),
			},
		},
		&armcdn.AFDOriginGroupsClientBeginCreateOptions{ResumeToken: ""})
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
