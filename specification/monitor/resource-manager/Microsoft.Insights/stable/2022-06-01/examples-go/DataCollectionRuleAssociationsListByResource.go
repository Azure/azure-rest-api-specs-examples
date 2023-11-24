package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionRuleAssociationsListByResource.json
func ExampleDataCollectionRuleAssociationsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.DataCollectionRuleAssociationProxyOnlyResourceListResult = armmonitor.DataCollectionRuleAssociationProxyOnlyResourceListResult{
		// 	Value: []*armmonitor.DataCollectionRuleAssociationProxyOnlyResource{
		// 		{
		// 			Name: to.Ptr("myRuleAssociation"),
		// 			Type: to.Ptr("Microsoft.Insights/dataCollectionRuleAssociations"),
		// 			Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
		// 			ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm/providers/Microsoft.Insights/dataCollectionRuleAssociations/myRuleAssociation"),
		// 			Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
		// 				DataCollectionRuleID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
		// 				ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleAssociationProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("herRuleAssociation"),
		// 			Type: to.Ptr("Microsoft.Insights/dataCollectionRuleAssociations"),
		// 			Etag: to.Ptr("3afa167b-3255-432b-b66d-e74a348468af"),
		// 			ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm/providers/Microsoft.Insights/dataCollectionRuleAssociations/herRuleAssociation"),
		// 			Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
		// 				DataCollectionRuleID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/herCollectionRule"),
		// 				ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleAssociationProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myEndpointAssociation"),
		// 			Type: to.Ptr("Microsoft.Insights/dataCollectionRuleAssociations"),
		// 			Etag: to.Ptr("562d96b1-29e9-4250-b2fd-8bebfdf77a9d"),
		// 			ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVm/providers/Microsoft.Insights/dataCollectionRuleAssociations/myEndpointAssociation"),
		// 			Properties: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceProperties{
		// 				DataCollectionEndpointID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myCollectionEndpoint"),
		// 				ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleAssociationProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armmonitor.DataCollectionRuleAssociationProxyOnlyResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
