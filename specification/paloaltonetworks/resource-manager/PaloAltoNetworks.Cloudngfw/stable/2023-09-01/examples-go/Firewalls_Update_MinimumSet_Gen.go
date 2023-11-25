package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/Firewalls_Update_MinimumSet_Gen.json
func ExampleFirewallsClient_Update_firewallsUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallsClient().Update(ctx, "firewall-rg", "firewall1", armpanngfw.FirewallResourceUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FirewallResource = armpanngfw.FirewallResource{
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpanngfw.FirewallDeploymentProperties{
	// 		DNSSettings: &armpanngfw.DNSSettings{
	// 		},
	// 		MarketplaceDetails: &armpanngfw.MarketplaceDetails{
	// 			OfferID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			PublisherID: to.Ptr("aaaa"),
	// 		},
	// 		NetworkProfile: &armpanngfw.NetworkProfile{
	// 			EnableEgressNat: to.Ptr(armpanngfw.EgressNatENABLED),
	// 			NetworkType: to.Ptr(armpanngfw.NetworkTypeVNET),
	// 			PublicIPs: []*armpanngfw.IPAddress{
	// 				{
	// 					Address: to.Ptr("20.22.92.11"),
	// 					ResourceID: to.Ptr("/subscriptions/01c7d41f-afaf-464e-8a8b-5c6f9f98cee8/resourceGroups/mj-liftr-integration/providers/Microsoft.Network/publicIPAddresses/mj-liftr-integration-PublicIp1"),
	// 			}},
	// 		},
	// 		PlanData: &armpanngfw.PlanData{
	// 			BillingCycle: to.Ptr(armpanngfw.BillingCycleWEEKLY),
	// 			PlanID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 	},
	// }
}
