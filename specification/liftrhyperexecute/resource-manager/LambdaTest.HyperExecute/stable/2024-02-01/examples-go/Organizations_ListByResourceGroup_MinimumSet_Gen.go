package armlambdatesthyperexecute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/lambdatesthyperexecute/armlambdatesthyperexecute"
)

// Generated from example definition: 2024-02-01/Organizations_ListByResourceGroup_MinimumSet_Gen.json
func ExampleOrganizationsClient_NewListByResourceGroupPager_organizationsListByResourceGroupMaximumSetGenGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlambdatesthyperexecute.NewClientFactory("171E7A75-341B-4472-BC4C-7603C5AB9F32", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationsClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armlambdatesthyperexecute.OrganizationsClientListByResourceGroupResponse{
		// 	OrganizationResourceListResult: armlambdatesthyperexecute.OrganizationResourceListResult{
		// 		Value: []*armlambdatesthyperexecute.OrganizationResource{
		// 			{
		// 				Location: to.Ptr("cvymsrlt"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Lambdatest.Hyoerexecute/organizations/testorg"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
