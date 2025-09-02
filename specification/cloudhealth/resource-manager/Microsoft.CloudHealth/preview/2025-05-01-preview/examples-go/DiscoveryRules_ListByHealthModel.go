package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/DiscoveryRules_ListByHealthModel.json
func ExampleDiscoveryRulesClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscoveryRulesClient().NewListByHealthModelPager("my-resource-group", "my-health-model", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcloudhealth.DiscoveryRulesClientListByHealthModelResponse{
		// 	DiscoveryRuleListResult: armcloudhealth.DiscoveryRuleListResult{
		// 		Value: []*armcloudhealth.DiscoveryRule{
		// 			{
		// 				Properties: &armcloudhealth.DiscoveryRuleProperties{
		// 					AuthenticationSetting: to.Ptr("authSetting1"),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("myDisplayName"),
		// 					DiscoverRelationships: to.Ptr(armcloudhealth.DiscoveryRuleRelationshipDiscoveryBehaviorEnabled),
		// 					AddRecommendedSignals: to.Ptr(armcloudhealth.DiscoveryRuleRecommendedSignalsBehaviorEnabled),
		// 					NumberOfDiscoveredEntities: to.Ptr[int32](5),
		// 					ResourceGraphQuery: to.Ptr("resources | where subscriptionId == '7ddfffd7-9b32-40df-1234-828cbd55d6f4' | where resourceGroup == 'my-rg'"),
		// 					EntityName: to.Ptr("f1f0ef5e-a5b5-4d02-b69c-7145f4658829"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.CloudHealth/healthModels/myHealthModel/discoveryRules/myDiscoveryRule"),
		// 				Name: to.Ptr("myDiscoveryRule"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthModels/discoveryRules"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("myCreatedBy"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("myLastModifiedBy"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ahgxpg"),
		// 	},
		// }
	}
}
