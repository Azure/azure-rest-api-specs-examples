package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListCdnPeeringPrefixes.json
func ExampleCdnPeeringPrefixesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCdnPeeringPrefixesClient().NewListPager("peeringLocation0", nil)
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
		// page.CdnPeeringPrefixListResult = armpeering.CdnPeeringPrefixListResult{
		// 	Value: []*armpeering.CdnPeeringPrefix{
		// 		{
		// 			Name: to.Ptr("CdnPrefix_192_168_1_0_24"),
		// 			Type: to.Ptr("Microsoft.Peering/cdnPeeringPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/cdnPeeringPrefixes/CdnPrefix_192_168_1_0_24"),
		// 			Properties: &armpeering.CdnPeeringPrefixProperties{
		// 				AzureRegion: to.Ptr("West Central US"),
		// 				AzureService: to.Ptr("AzureCompute"),
		// 				BgpCommunity: to.Ptr("0000:0000"),
		// 				IsPrimaryRegion: to.Ptr(false),
		// 				Prefix: to.Ptr("192.168.1.0/24"),
		// 			},
		// 	}},
		// }
	}
}
