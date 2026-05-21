package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth/v2"
)

// Generated from example definition: 2025-05-01/ImpactedResources_ListByTenantId_ListByEventId.json
func ExampleImpactedResourcesClient_NewListByTenantIDAndEventIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImpactedResourcesClient().NewListByTenantIDAndEventIDPager("BC_1-FXZ", &armresourcehealth.ImpactedResourcesClientListByTenantIDAndEventIDOptions{
		Filter: to.Ptr("targetRegion eq 'westus'")})
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
		// page = armresourcehealth.ImpactedResourcesClientListByTenantIDAndEventIDResponse{
		// 	EventImpactedResourceListResult: armresourcehealth.EventImpactedResourceListResult{
		// 		NextLink: to.Ptr("{originalRequestUrl}?$skipToken={OpaquePageNumber}"),
		// 		Value: []*armresourcehealth.EventImpactedResource{
		// 			{
		// 				Name: to.Ptr("abc-123-ghj-456"),
		// 				Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 				ID: to.Ptr("/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/abc-123-ghj-456"),
		// 				Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 					TargetRegion: to.Ptr("westus"),
		// 					TargetResourceID: to.Ptr("{resourceId-1}"),
		// 					TargetResourceType: to.Ptr("Microsoft.Compute/VirtualMachines"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("wer-345-tyu-789"),
		// 				Type: to.Ptr("Microsoft.ResourceHealth/events/impactedResources"),
		// 				ID: to.Ptr("/providers/Microsoft.ResourceHealth/events/BC_1-FXZ/impactedResources/wer-345-tyu-789"),
		// 				Properties: &armresourcehealth.EventImpactedResourceProperties{
		// 					TargetRegion: to.Ptr("westus"),
		// 					TargetResourceID: to.Ptr("{resourceId-2}"),
		// 					TargetResourceType: to.Ptr("Microsoft.Compute/VirtualMachines"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
