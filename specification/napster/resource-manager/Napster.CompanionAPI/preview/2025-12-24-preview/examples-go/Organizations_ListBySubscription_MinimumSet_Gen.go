package armnapsteromniagentapi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/napsteromniagentapi/armnapsteromniagentapi"
)

// Generated from example definition: 2025-12-24-preview/Organizations_ListBySubscription_MinimumSet_Gen.json
func ExampleOrganizationsClient_NewListBySubscriptionPager_organizationsListBySubscriptionMinimumSetGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnapsteromniagentapi.NewClientFactory("0F0FBCF9-8374-47FC-B189-B79B84033EA3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationsClient().NewListBySubscriptionPager(nil)
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
		// page = armnapsteromniagentapi.OrganizationsClientListBySubscriptionResponse{
		// 	OrganizationResourceListResult: armnapsteromniagentapi.OrganizationResourceListResult{
		// 		Value: []*armnapsteromniagentapi.OrganizationResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/0F0FBCF9-8374-47FC-B189-B79B84033EA3/resourceGroups/rgopenapi/providers/Napster.CompanionAPI/organizations/contosoOrg"),
		// 				Location: to.Ptr("eastus"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
