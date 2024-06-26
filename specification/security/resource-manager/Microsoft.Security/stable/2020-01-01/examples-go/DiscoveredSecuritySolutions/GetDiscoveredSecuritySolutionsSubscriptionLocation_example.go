package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions/GetDiscoveredSecuritySolutionsSubscriptionLocation_example.json
func ExampleDiscoveredSecuritySolutionsClient_NewListByHomeRegionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscoveredSecuritySolutionsClient().NewListByHomeRegionPager("centralus", nil)
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
		// page.DiscoveredSecuritySolutionList = armsecurity.DiscoveredSecuritySolutionList{
		// 	Value: []*armsecurity.DiscoveredSecuritySolution{
		// 		{
		// 			Location: to.Ptr("eastus"),
		// 			Name: to.Ptr("CP"),
		// 			Type: to.Ptr("Microsoft.Security/locations/discoveredSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Security/locations/centralus/discoveredSecuritySolutions/CP"),
		// 			Properties: &armsecurity.DiscoveredSecuritySolutionProperties{
		// 				Offer: to.Ptr("cisco-asav"),
		// 				Publisher: to.Ptr("cisco"),
		// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilyNgfw),
		// 				SKU: to.Ptr("asav-azure-byol"),
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("eastus2"),
		// 			Name: to.Ptr("paloalto7"),
		// 			Type: to.Ptr("Microsoft.Security/locations/discoveredSecuritySolutions"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg2/providers/Microsoft.Security/locations/centralus/discoveredSecuritySolutions/paloalto7"),
		// 			Properties: &armsecurity.DiscoveredSecuritySolutionProperties{
		// 				Offer: to.Ptr("vmseries1"),
		// 				Publisher: to.Ptr("paloaltonetworks"),
		// 				SecurityFamily: to.Ptr(armsecurity.SecurityFamilyNgfw),
		// 				SKU: to.Ptr("byol"),
		// 			},
		// 	}},
		// }
	}
}
