package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Usages_ListByLocation.json
func ExampleUsagesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListByLocationPager("westus", nil)
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
		// page.ListUsagesResult = armdevcenter.ListUsagesResult{
		// 	Value: []*armdevcenter.Usage{
		// 		{
		// 			Name: &armdevcenter.UsageName{
		// 				Value: to.Ptr("devcenters"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](2),
		// 			Limit: to.Ptr[int64](8),
		// 			Unit: to.Ptr(armdevcenter.UsageUnitCount),
		// 		},
		// 		{
		// 			Name: &armdevcenter.UsageName{
		// 				Value: to.Ptr("projects"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](5),
		// 			Limit: to.Ptr[int64](30),
		// 			Unit: to.Ptr(armdevcenter.UsageUnitCount),
		// 	}},
		// }
	}
}
