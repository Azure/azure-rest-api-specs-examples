```go
package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateJitRequest.json
func ExampleJitRequestsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedapplications.NewJitRequestsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg",
		"myJitRequest",
		armmanagedapplications.JitRequestDefinition{
			Properties: &armmanagedapplications.JitRequestProperties{
				ApplicationResourceID: to.Ptr("/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158"),
				JitAuthorizationPolicies: []*armmanagedapplications.JitAuthorizationPolicies{
					{
						PrincipalID:      to.Ptr("1db8e132e2934dbcb8e1178a61319491"),
						RoleDefinitionID: to.Ptr("ecd05a23-931a-4c38-a52b-ac7c4c583334"),
					}},
				JitSchedulingPolicy: &armmanagedapplications.JitSchedulingPolicy{
					Type:      to.Ptr(armmanagedapplications.JitSchedulingTypeOnce),
					Duration:  to.Ptr("PT8H"),
					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-22T05:48:30.6661804Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsolutions%2Farmmanagedapplications%2Fv1.0.0/sdk/resourcemanager/solutions/armmanagedapplications/README.md) on how to add the SDK to your project and authenticate.
