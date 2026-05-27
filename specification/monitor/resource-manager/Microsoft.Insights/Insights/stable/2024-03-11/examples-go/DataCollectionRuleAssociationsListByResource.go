package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-11/DataCollectionRuleAssociationsListByResource.json
func ExampleDataCollectionRuleAssociationsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataCollectionRuleAssociationsClient().NewListByResourcePager("subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm", nil)
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
		// page = armmonitor.DataCollectionRuleAssociationsClientListByResourceResponse{
		// 	DataCollectionRuleAssociationProxyOnlyResourceListResult: armmonitor.DataCollectionRuleAssociationProxyOnlyResourceListResult{
		// 		Value: []*armmonitor.DataCollectionRuleAssociationProxyOnlyResource{
		// 			{
		// 				Name: to.Ptr("myAssociation"),
		// 				Type: to.Ptr("Microsoft.Insights/dataCollectionRuleAssociations"),
		// 				Etag: to.Ptr("00000836-0000-0100-0000-65fdd8bc0000"),
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm/providers/Microsoft.Insights/dataCollectionRuleAssociations/myAssociation"),
		// 				Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
		// 					DataCollectionRuleID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
		// 				},
		// 				SystemData: &armmonitor.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T19:15:08.2155156Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T19:15:08.2155156Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user1"),
		// 					LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("configurationAccessEndpoint"),
		// 				Type: to.Ptr("Microsoft.Insights/dataCollectionRuleAssociations"),
		// 				Etag: to.Ptr("00000836-0000-0100-0000-65fdd8bc0000"),
		// 				ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm/providers/Microsoft.Insights/dataCollectionRuleAssociations/configurationAccessEndpoint"),
		// 				Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
		// 					DataCollectionEndpointID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myCollectionEndpoint"),
		// 				},
		// 				SystemData: &armmonitor.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T19:15:08.2155156Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T19:15:08.2155156Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user1"),
		// 					LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
