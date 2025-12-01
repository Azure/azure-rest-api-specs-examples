package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/MetricsObjectFirewall_ListByFirewalls_MaximumSet_Gen.json
func ExampleMetricsObjectFirewallClient_NewListByFirewallsPager_metricsObjectFirewallListByFirewallsMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("aaaaaaaaaaaaaaaaaaaaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetricsObjectFirewallClient().NewListByFirewallsPager("rgopenapi", "IFTDk", nil)
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
		// page = armpanngfw.MetricsObjectFirewallClientListByFirewallsResponse{
		// 	MetricsObjectFirewallResourceListResult: armpanngfw.MetricsObjectFirewallResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/aaaaaaaaaaaaaaaaaaaaaaaaa/resourceGroups/rgopenapi/providers/PaloAltoNetworks.Cloudngfw/firewalls/firewall1/metrics?api-version=2025-10-08&$skiptoken=xyz"),
		// 		Value: []*armpanngfw.MetricsObjectFirewallResource{
		// 			{
		// 				Name: to.Ptr("aaaaaa"),
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				ID: to.Ptr("aaaaaaa"),
		// 				Properties: &armpanngfw.MetricsObject{
		// 					ApplicationInsightsConnectionString: to.Ptr("aaaaaaaaa"),
		// 					ApplicationInsightsResourceID: to.Ptr("aaaaaaaaaa"),
		// 					PanEtag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
		// 				},
		// 				SystemData: &armpanngfw.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-17T09:12:09.452Z"); return t}()),
		// 					CreatedBy: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-17T09:12:09.452Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
