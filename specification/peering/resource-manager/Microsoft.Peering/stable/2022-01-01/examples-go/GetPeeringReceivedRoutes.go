package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/GetPeeringReceivedRoutes.json
func ExampleReceivedRoutesClient_NewListByPeeringPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReceivedRoutesClient().NewListByPeeringPager("rgName", "peeringName", &armpeering.ReceivedRoutesClientListByPeeringOptions{Prefix: to.Ptr("1.1.1.0/24"),
		AsPath:                  to.Ptr("123 456"),
		OriginAsValidationState: to.Ptr("Valid"),
		RpkiValidationState:     to.Ptr("Valid"),
		SkipToken:               nil,
	})
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
		// page.ReceivedRouteListResult = armpeering.ReceivedRouteListResult{
		// 	Value: []*armpeering.ReceivedRoute{
		// 		{
		// 			AsPath: to.Ptr("123 456"),
		// 			NextHop: to.Ptr("127.0.0.1"),
		// 			OriginAsValidationState: to.Ptr("Valid"),
		// 			Prefix: to.Ptr("1.1.1.0/24"),
		// 			ReceivedTimestamp: to.Ptr("2020-04-05 03:39:20"),
		// 			RpkiValidationState: to.Ptr("Valid"),
		// 			TrustAnchor: to.Ptr("Arin"),
		// 	}},
		// }
	}
}
