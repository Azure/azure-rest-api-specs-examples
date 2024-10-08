package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/56537883b7cdb95618c3d1ec1c0ee37b59d88d72/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_ListBySubscription_MinimumSet_Gen.json
func ExampleOrganizationsClient_NewListBySubscriptionPager_organizationsListBySubscriptionMin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.InformaticaOrganizationResourceListResult = arminformaticadatamgmt.InformaticaOrganizationResourceListResult{
		// 	Value: []*arminformaticadatamgmt.InformaticaOrganizationResource{
		// 		{
		// 			ID: to.Ptr("cadokiejnrth"),
		// 			Location: to.Ptr("pamjoudtssthlbhrnfjidr"),
		// 	}},
		// }
	}
}
