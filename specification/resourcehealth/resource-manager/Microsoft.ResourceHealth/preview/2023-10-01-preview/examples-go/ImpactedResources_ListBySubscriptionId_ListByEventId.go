package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ImpactedResources_ListBySubscriptionId_ListByEventId.json
func ExampleImpactedResourcesClient_NewListBySubscriptionIDAndEventIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImpactedResourcesClient().NewListBySubscriptionIDAndEventIDPager("BC_1-FXZ", &armresourcehealth.ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions{Filter: to.Ptr("targetRegion eq 'westus'")})
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
		// page.EventImpactedResourceListResult = armresourcehealth.EventImpactedResourceListResult{
		// 	Value: []*armresourcehealth.EventImpactedResource{
		// 		{
		// 			Name: to.Ptr("abc-123-ghj-456"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 			ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/abc-123-ghj-456"),
		// 			Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 				MaintenanceEndTime: to.Ptr("2023-08-30T00:00:00.00Z"),
		// 				MaintenanceStartTime: to.Ptr("2023-08-15T23:32:39.76Z"),
		// 				ResourceGroup: to.Ptr("TEST"),
		// 				ResourceName: to.Ptr("testvm"),
		// 				Status: to.Ptr("Pending"),
		// 				TargetRegion: to.Ptr("westus"),
		// 				TargetResourceID: to.Ptr("/subscriptions/4970d23e-ed41-4670-9c19-02a1d2808dd9/resourceGroups/TEST/providers/Microsoft.Compute/virtualMachines/testvm"),
		// 				TargetResourceType: to.Ptr("Microsoft.Compute/VirtualMachines"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("jkl-901-hgy-445"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 			ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/jkl-901-hgy-445"),
		// 			Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 				Info: []*armresourcehealth.KeyValueItem{
		// 					{
		// 						Key: to.Ptr("ContainerGrp"),
		// 						Value: to.Ptr("xyz"),
		// 					},
		// 					{
		// 						Key: to.Ptr("UserGuid"),
		// 						Value: to.Ptr("guid"),
		// 					},
		// 					{
		// 						Key: to.Ptr("AplicationID"),
		// 						Value: to.Ptr("Abc"),
		// 					},
		// 					{
		// 						Key: to.Ptr("UserPrincipalObjectId"),
		// 						Value: to.Ptr("def"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("wer-345-tyu-789"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 			ID: to.Ptr("/subscriptions/{subscripitionId}/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/wer-345-tyu-789"),
		// 			Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 				MaintenanceEndTime: to.Ptr("2023-08-30T00:00:00.00Z"),
		// 				MaintenanceStartTime: to.Ptr("2023-08-15T23:32:39.76Z"),
		// 				ResourceGroup: to.Ptr("Dev2"),
		// 				ResourceName: to.Ptr("testvm2"),
		// 				Status: to.Ptr("Pending"),
		// 				TargetRegion: to.Ptr("westus"),
		// 				TargetResourceID: to.Ptr("/subscriptions/4970d23e-ed41-4670-9c19-02a1d2808dd9/resourceGroups/Dev2/providers/Microsoft.Compute/virtualMachines/testvm2"),
		// 				TargetResourceType: to.Ptr("Microsoft.Compute/VirtualMachines"),
		// 			},
		// 	}},
		// }
	}
}
