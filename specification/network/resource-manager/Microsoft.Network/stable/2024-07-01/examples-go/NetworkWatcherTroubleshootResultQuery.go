package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkWatcherTroubleshootResultQuery.json
func ExampleWatchersClient_BeginGetTroubleshootingResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginGetTroubleshootingResult(ctx, "rg1", "nw1", armnetwork.QueryTroubleshootingParameters{
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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
	// res.TroubleshootingResult = armnetwork.TroubleshootingResult{
	// 	Code: to.Ptr("UnHealthy"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T00:20:09.914Z"); return t}()),
	// 	Results: []*armnetwork.TroubleshootingDetails{
	// 		{
	// 			Detail: to.Ptr("During this time S2S VPN tunnels to on premises sites or other Azure virtual networks will be disconnected"),
	// 			ID: to.Ptr("000000"),
	// 			ReasonType: to.Ptr("VipUnResponsive"),
	// 			RecommendedActions: []*armnetwork.TroubleshootingRecommendedActions{
	// 				{
	// 					ActionText: to.Ptr("Verify if there is a network security group (NSG) applied to the GatewaySubnet"),
	// 					ActionURI: to.Ptr("https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-create-nsg-arm-pportal"),
	// 					ActionURIText: to.Ptr("Verify"),
	// 				},
	// 				{
	// 					ActionText: to.Ptr("If your VPN gateway isn't up and running by the expected resolution time, contact support"),
	// 					ActionURI: to.Ptr("http://azure.microsoft.com/support"),
	// 					ActionURIText: to.Ptr("contact support"),
	// 			}},
	// 			Summary: to.Ptr("We are sorry, your VPN gateway is unreachable from the Internet"),
	// 	}},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T00:19:47.044Z"); return t}()),
	// }
}
