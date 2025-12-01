package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/Firewalls_ListBySubscription_MinimumSet_Gen.json
func ExampleFirewallsClient_NewListBySubscriptionPager_firewallsListBySubscriptionMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallsClient().NewListBySubscriptionPager(nil)
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
		// page = armpanngfw.FirewallsClientListBySubscriptionResponse{
		// 	FirewallResourceListResult: armpanngfw.FirewallResourceListResult{
		// 		Value: []*armpanngfw.FirewallResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/providers/PaloAltoNetworks.Cloudngfw/firewalls/firewall"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armpanngfw.FirewallDeploymentProperties{
		// 					DNSSettings: &armpanngfw.DNSSettings{
		// 					},
		// 					MarketplaceDetails: &armpanngfw.MarketplaceDetails{
		// 						OfferID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						PublisherID: to.Ptr("aaaa"),
		// 					},
		// 					NetworkProfile: &armpanngfw.NetworkProfile{
		// 						EnableEgressNat: to.Ptr(armpanngfw.EgressNatENABLED),
		// 						NetworkType: to.Ptr(armpanngfw.NetworkTypeVNET),
		// 						PublicIPs: []*armpanngfw.IPAddress{
		// 							{
		// 								Address: to.Ptr("20.22.92.11"),
		// 								ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
		// 							},
		// 						},
		// 					},
		// 					PlanData: &armpanngfw.PlanData{
		// 						BillingCycle: to.Ptr(armpanngfw.BillingCycleWEEKLY),
		// 						PlanID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
