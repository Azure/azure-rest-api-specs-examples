package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2027-01-01/PublicCloudConnectors_ListBySubscription_MinimumSet_Gen.json
func ExamplePublicCloudConnectorsClient_NewListBySubscriptionPager_publicCloudConnectorsListBySubscriptionGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicCloudConnectorsClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").NewListBySubscriptionPager(nil)
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
		// page = armhybridconnectivity.PublicCloudConnectorsClientListBySubscriptionResponse{
		// 	PublicCloudConnectorListResult: armhybridconnectivity.PublicCloudConnectorListResult{
		// 		Value: []*armhybridconnectivity.PublicCloudConnector{
		// 			{
		// 				ID: to.Ptr("/subscriptions/5ACC4579-DB34-4C2F-8F8C-25061168F342/providers/Microsoft.HybridConnectivity/PublicCloudConnectors/esixipkbydb"),
		// 				Location: to.Ptr("jpiglusfxynfcewcjwvvnn"),
		// 				Kind: to.Ptr(armhybridconnectivity.HostTypeGCP),
		// 			},
		// 		},
		// 	},
		// }
	}
}
