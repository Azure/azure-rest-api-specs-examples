package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fdf43f2fdacf17bd78c0621df44a5c024b61db82/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FirewallStatus_ListByFirewalls_MaximumSet_Gen.json
func ExampleFirewallStatusClient_NewListByFirewallsPager_firewallStatusListByFirewallsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirewallStatusClient().NewListByFirewallsPager("rgopenapi", "firewall1", nil)
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
		// page.FirewallStatusResourceListResult = armpanngfw.FirewallStatusResourceListResult{
		// 	Value: []*armpanngfw.FirewallStatusResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("aaaaaaa"),
		// 			ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/firewall-rg/providers/PaloAltoNetworks.Cloudngfw/firewalls/firewall1/statuses/default"),
		// 			SystemData: &armpanngfw.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				CreatedBy: to.Ptr("praval"),
		// 				CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("praval"),
		// 				LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 			},
		// 			Properties: &armpanngfw.FirewallStatusProperty{
		// 				HealthReason: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				HealthStatus: to.Ptr(armpanngfw.HealthStatusGREEN),
		// 				IsPanoramaManaged: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 				PanoramaStatus: &armpanngfw.PanoramaStatus{
		// 					PanoramaServer2Status: to.Ptr(armpanngfw.ServerStatusUP),
		// 					PanoramaServerStatus: to.Ptr(armpanngfw.ServerStatusUP),
		// 				},
		// 				ProvisioningState: to.Ptr(armpanngfw.ReadOnlyProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
