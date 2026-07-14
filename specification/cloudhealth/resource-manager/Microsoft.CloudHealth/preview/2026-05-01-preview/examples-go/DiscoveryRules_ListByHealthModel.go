package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/DiscoveryRules_ListByHealthModel.json
func ExampleDiscoveryRulesClient_NewListByHealthModelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscoveryRulesClient().NewListByHealthModelPager("online-store-rg", "online-store", nil)
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
		// 					AuthenticationSetting: to.Ptr("default-auth"),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("Discover web apps"),
		// 					DiscoverRelationships: to.Ptr(armcloudhealth.DiscoveryRuleRelationshipDiscoveryBehaviorEnabled),
		// 					AddRecommendedSignals: to.Ptr(armcloudhealth.DiscoveryRuleRecommendedSignalsBehaviorEnabled),
		// 					Specification: &armcloudhealth.ResourceGraphQuerySpecification{
		// 						Kind: to.Ptr(armcloudhealth.DiscoveryRuleKindResourceGraphQuery),
		// 						ResourceGraphQuery: to.Ptr("resources | where type =~ 'microsoft.web/sites' and resourceGroup =~ 'online-store-rg' | project id, name, location"),
		// 					},
		// 					AddResourceHealthSignal: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorEnabled),
		// 					EntityName: to.Ptr("online-store-web-apps"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/discoveryRules/discover-web-apps"),
		// 				Name: to.Ptr("discover-web-apps"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/discoveryRules"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcloudhealth.DiscoveryRuleProperties{
		// 					AuthenticationSetting: to.Ptr("default-auth"),
		// 					ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
		// 					DisplayName: to.Ptr("Discover Application Insights topology"),
		// 					DiscoverRelationships: to.Ptr(armcloudhealth.DiscoveryRuleRelationshipDiscoveryBehaviorEnabled),
		// 					AddRecommendedSignals: to.Ptr(armcloudhealth.DiscoveryRuleRecommendedSignalsBehaviorEnabled),
		// 					Specification: &armcloudhealth.ApplicationInsightsTopologySpecification{
		// 						Kind: to.Ptr(armcloudhealth.DiscoveryRuleKindApplicationInsightsTopology),
		// 						ApplicationInsightsResourceID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.Insights/components/online-store-ai"),
		// 					},
		// 					AddResourceHealthSignal: to.Ptr(armcloudhealth.ResourceHealthAvailabilityStateSignalBehaviorDisabled),
		// 					EntityName: to.Ptr("online-store-topology"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/discoveryRules/discover-app-topology"),
		// 				Name: to.Ptr("discover-app-topology"),
		// 				Type: to.Ptr("Microsoft.CloudHealth/healthmodels/discoveryRules"),
		// 				SystemData: &armcloudhealth.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
