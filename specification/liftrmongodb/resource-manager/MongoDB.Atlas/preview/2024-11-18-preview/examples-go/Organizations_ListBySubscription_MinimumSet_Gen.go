package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2024-11-18-preview/Organizations_ListBySubscription_MinimumSet_Gen.json
func ExampleOrganizationsClient_NewListBySubscriptionPager_organizationsListBySubscriptionMaximumSetGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("422A4D59-A5BC-4DBB-8831-EC666633F64F", cred, nil)
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
		// page = armmongodbatlas.OrganizationsClientListBySubscriptionResponse{
		// 	OrganizationResourceListResult: armmongodbatlas.OrganizationResourceListResult{
		// 		Value: []*armmongodbatlas.OrganizationResource{
		// 			{
		// 				Location: to.Ptr("tjmlissxvwsk"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/MongoDB.Atlas/organizations/testorg"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
