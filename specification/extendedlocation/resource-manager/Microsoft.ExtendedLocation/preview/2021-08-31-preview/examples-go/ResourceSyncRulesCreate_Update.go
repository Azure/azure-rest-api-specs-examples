package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fb9c8e2ca33e9723c2b2fc849f627329067feb54/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/ResourceSyncRulesCreate_Update.json
func ExampleResourceSyncRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceSyncRulesClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "customLocation01", "resourceSyncRule01", armextendedlocation.ResourceSyncRule{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceSyncRule = armextendedlocation.ResourceSyncRule{
	// 	Name: to.Ptr("resourceSyncRule01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations/resourceSyncRules"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/customLocations/customLocation01/resourcesyncrules/resourceSyncRule01"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armextendedlocation.ResourceSyncRuleProperties{
	// 		Priority: to.Ptr[int32](999),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Selector: &armextendedlocation.ResourceSyncRulePropertiesSelector{
	// 			MatchExpressions: []*armextendedlocation.MatchExpressionsProperties{
	// 				{
	// 					Key: to.Ptr("key4"),
	// 					Operator: to.Ptr("In"),
	// 					Values: []*string{
	// 						to.Ptr("value4")},
	// 				}},
	// 				MatchLabels: map[string]*string{
	// 					"key1": to.Ptr("value1"),
	// 				},
	// 			},
	// 			TargetResourceGroup: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-15T18:53:29.0928001Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-15T18:53:29.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}
