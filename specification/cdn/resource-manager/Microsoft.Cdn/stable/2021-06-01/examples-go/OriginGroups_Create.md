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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Create.json
func ExampleOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewOriginGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		"<origin-group-name>",
		armcdn.OriginGroup{
			Properties: &armcdn.OriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Ptr[int32](120),
					ProbePath:              to.Ptr("<probe-path>"),
					ProbeProtocol:          to.Ptr(armcdn.ProbeProtocolHTTP),
					ProbeRequestType:       to.Ptr(armcdn.HealthProbeRequestTypeGET),
				},
				Origins: []*armcdn.ResourceReference{
					{
						ID: to.Ptr("<id>"),
					}},
				ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
					ResponseBasedDetectedErrorTypes:          to.Ptr(armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
					ResponseBasedFailoverThresholdPercentage: to.Ptr[int32](10),
				},
			},
		},
		&armcdn.OriginGroupsClientBeginCreateOptions{ResumeToken: ""})
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
