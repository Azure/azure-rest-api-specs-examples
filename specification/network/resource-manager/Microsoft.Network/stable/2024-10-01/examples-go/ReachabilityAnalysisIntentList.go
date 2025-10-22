package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/ReachabilityAnalysisIntentList.json
func ExampleReachabilityAnalysisIntentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReachabilityAnalysisIntentsClient().NewListPager("rg1", "testNetworkManager", "testVerifierWorkspace1", &armnetwork.ReachabilityAnalysisIntentsClientListOptions{SkipToken: nil,
		Skip:      nil,
		Top:       nil,
		SortKey:   nil,
		SortValue: nil,
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
		// page.ReachabilityAnalysisIntentListResult = armnetwork.ReachabilityAnalysisIntentListResult{
		// 	Value: []*armnetwork.ReachabilityAnalysisIntent{
		// 		{
		// 			Name: to.Ptr("testReachabilityAnalysisIntenant1"),
		// 			Type: to.Ptr("Microsoft.Network/networkManagers/verifierWorkspaces/reachabilityAnalysisIntents"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisIntents/testReachabilityAnalysisIntenant1"),
		// 			SystemData: &armnetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetwork.ReachabilityAnalysisIntentProperties{
		// 				Description: to.Ptr("A sample  reachability analysis intent"),
		// 				DestinationResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmDest"),
		// 				IPTraffic: &armnetwork.IPTraffic{
		// 					DestinationIPs: []*string{
		// 						to.Ptr("10.4.0.1")},
		// 						DestinationPorts: []*string{
		// 							to.Ptr("0")},
		// 							Protocols: []*armnetwork.NetworkProtocol{
		// 								to.Ptr(armnetwork.NetworkProtocolAny)},
		// 								SourceIPs: []*string{
		// 									to.Ptr("10.4.0.0")},
		// 									SourcePorts: []*string{
		// 										to.Ptr("0")},
		// 									},
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmSrc"),
		// 								},
		// 						}},
		// 					}
	}
}
