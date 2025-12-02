package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/MetricsObjectFirewall_CreateOrUpdate_MaximumSet_Gen.json
func ExampleMetricsObjectFirewallClient_BeginCreateOrUpdate_metricsObjectFirewallCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("aaaaaaa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMetricsObjectFirewallClient().BeginCreateOrUpdate(ctx, "rgopenapi", "aaaaaaaaaaaaaaaaaaaaaaaa", armpanngfw.MetricsObjectFirewallResource{
		Properties: &armpanngfw.MetricsObject{
			ApplicationInsightsConnectionString: to.Ptr("aaa"),
			ApplicationInsightsResourceID:       to.Ptr("aaaaaaaaaaaaaaa"),
			PanEtag:                             to.Ptr("aaaaaaaaaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.MetricsObjectFirewallClientCreateOrUpdateResponse{
	// 	MetricsObjectFirewallResource: &armpanngfw.MetricsObjectFirewallResource{
	// 		Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 		ID: to.Ptr("aaaaaaaaaaaa"),
	// 		Properties: &armpanngfw.MetricsObject{
	// 			ApplicationInsightsConnectionString: to.Ptr("aaa"),
	// 			ApplicationInsightsResourceID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			PanEtag: to.Ptr("aaaaaaaaaa"),
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
	// 		},
	// 		SystemData: &armpanngfw.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-04T09:21:06.704Z"); return t}()),
	// 			CreatedBy: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-04T09:21:06.704Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("aaaaaaaaaaa"),
	// 			LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
