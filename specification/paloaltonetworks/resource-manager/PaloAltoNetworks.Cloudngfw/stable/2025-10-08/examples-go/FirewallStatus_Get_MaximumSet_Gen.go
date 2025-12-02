package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/FirewallStatus_Get_MaximumSet_Gen.json
func ExampleFirewallStatusClient_Get_firewallStatusGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallStatusClient().Get(ctx, "rgopenapi", "firewall1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.FirewallStatusClientGetResponse{
	// 	FirewallStatusResource: &armpanngfw.FirewallStatusResource{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("aaaa"),
	// 		ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourcegroups/firewall-rg/providers/PaloAltoNetworks.Cloudngfw/firewalls/firewall1/statuses/default"),
	// 		Properties: &armpanngfw.FirewallStatusProperty{
	// 			HealthReason: to.Ptr("aaaaaaaaaaaa"),
	// 			HealthStatus: to.Ptr(armpanngfw.HealthStatusGREEN),
	// 			IsPanoramaManaged: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 			PanoramaStatus: &armpanngfw.PanoramaStatus{
	// 				PanoramaServer2Status: to.Ptr(armpanngfw.ServerStatusUP),
	// 				PanoramaServerStatus: to.Ptr(armpanngfw.ServerStatusUP),
	// 			},
	// 			ProvisioningState: to.Ptr(armpanngfw.ReadOnlyProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armpanngfw.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			CreatedBy: to.Ptr("praval"),
	// 			CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("praval"),
	// 			LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
