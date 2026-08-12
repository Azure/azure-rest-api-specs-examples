package armfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
)

// Generated from example definition: 2026-08-01-preview/FabricCapacities_ListUsagesBySubscription.json
func ExampleCapacitiesClient_NewListUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfabric.NewClientFactory("548B7FB7-3B2A-4F46-BB02-66473F1FC22C", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCapacitiesClient().NewListUsagesPager("centraluseuap", nil)
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
		// page = armfabric.CapacitiesClientListUsagesResponse{
		// 	PagedQuota: armfabric.PagedQuota{
		// 		Value: []*armfabric.Quota{
		// 			{
		// 				Name: &armfabric.QuotaName{
		// 					Value: to.Ptr("CapacityQuota"),
		// 					LocalizedValue: to.Ptr("CapacityQuota"),
		// 				},
		// 				Unit: to.Ptr("CU"),
		// 				CurrentValue: to.Ptr[int64](378),
		// 				Limit: to.Ptr[int64](512),
		// 			},
		// 		},
		// 	},
		// }
	}
}
