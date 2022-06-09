```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/ApplicationGroup/ApplicationGroupCreate.json
func ExampleApplicationGroupClient_CreateOrUpdateApplicationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewApplicationGroupClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateApplicationGroup(ctx,
		"contosotest",
		"contoso-ua-test-eh-system-1",
		"appGroup1",
		armeventhub.ApplicationGroup{
			Properties: &armeventhub.ApplicationGroupProperties{
				ClientAppGroupIdentifier: to.Ptr("SASKeyName=KeyName"),
				IsEnabled:                to.Ptr(true),
				Policies: []armeventhub.ApplicationGroupPolicyClassification{
					&armeventhub.ThrottlingPolicy{
						Name:               to.Ptr("ThrottlingPolicy1"),
						Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
						MetricID:           to.Ptr(armeventhub.MetricIDIncomingMessages),
						RateLimitThreshold: to.Ptr[int64](7912),
					},
					&armeventhub.ThrottlingPolicy{
						Name:               to.Ptr("ThrottlingPolicy2"),
						Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
						MetricID:           to.Ptr(armeventhub.MetricIDIncomingBytes),
						RateLimitThreshold: to.Ptr[int64](3951729),
					},
					&armeventhub.ThrottlingPolicy{
						Name:               to.Ptr("ThrottlingPolicy3"),
						Type:               to.Ptr(armeventhub.ApplicationGroupPolicyTypeThrottlingPolicy),
						MetricID:           to.Ptr(armeventhub.MetricIDOutgoingBytes),
						RateLimitThreshold: to.Ptr[int64](245175),
					}},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv1.1.0-beta.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.
