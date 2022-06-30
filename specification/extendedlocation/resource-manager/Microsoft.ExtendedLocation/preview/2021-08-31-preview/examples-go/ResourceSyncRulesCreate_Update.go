package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesCreate_Update.json
func ExampleResourceSyncRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armextendedlocation.NewResourceSyncRulesClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testresourcegroup",
		"customLocation01",
		"resourceSyncRule01",
		armextendedlocation.ResourceSyncRule{
			Location: to.Ptr("West US"),
			Properties: &armextendedlocation.ResourceSyncRuleProperties{
				Priority: to.Ptr[int32](999),
				Selector: &armextendedlocation.ResourceSyncRulePropertiesSelector{
					MatchExpressions: []*armextendedlocation.MatchExpressionsProperties{
						{
							Key:      to.Ptr("key4"),
							Operator: to.Ptr("In"),
							Values: []*string{
								to.Ptr("value4")},
						}},
					MatchLabels: map[string]*string{
						"key1": to.Ptr("value1"),
					},
				},
				TargetResourceGroup: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup"),
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
