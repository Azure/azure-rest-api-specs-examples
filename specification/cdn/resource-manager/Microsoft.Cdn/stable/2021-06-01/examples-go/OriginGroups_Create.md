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

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Create.json
func ExampleOriginGroupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewOriginGroupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<endpoint-name>",
		"<origin-group-name>",
		armcdn.OriginGroup{
			Properties: &armcdn.OriginGroupProperties{
				HealthProbeSettings: &armcdn.HealthProbeParameters{
					ProbeIntervalInSeconds: to.Int32Ptr(120),
					ProbePath:              to.StringPtr("<probe-path>"),
					ProbeProtocol:          armcdn.ProbeProtocolHTTP.ToPtr(),
					ProbeRequestType:       armcdn.HealthProbeRequestTypeGET.ToPtr(),
				},
				Origins: []*armcdn.ResourceReference{
					{
						ID: to.StringPtr("<id>"),
					}},
				ResponseBasedOriginErrorDetectionSettings: &armcdn.ResponseBasedOriginErrorDetectionParameters{
					ResponseBasedDetectedErrorTypes:          armcdn.ResponseBasedDetectedErrorTypesTCPErrorsOnly.ToPtr(),
					ResponseBasedFailoverThresholdPercentage: to.Int32Ptr(10),
				},
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
	log.Printf("Response result: %#v\n", res.OriginGroupsClientCreateResult)
}
```
