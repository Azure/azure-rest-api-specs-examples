package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ReachabilityAnalysisRunList.json
func ExampleReachabilityAnalysisRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReachabilityAnalysisRunsClient().NewListPager("rg1", "testNetworkManager", "testVerifierWorkspace1", &armnetwork.ReachabilityAnalysisRunsClientListOptions{SkipToken: nil,
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
		// page.ReachabilityAnalysisRunListResult = armnetwork.ReachabilityAnalysisRunListResult{
		// 	Value: []*armnetwork.ReachabilityAnalysisRun{
		// 		{
		// 			Name: to.Ptr("testReachabilityAnalysisRun1"),
		// 			Type: to.Ptr("Microsoft.Network/networkManagers/verifierWorkspaces/reachabilityAnalysisRuns"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisRuns/testReachabilityAnalysisRun1"),
		// 			SystemData: &armnetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetwork.ReachabilityAnalysisRunProperties{
		// 				Description: to.Ptr("A sample  reachability analysis intent"),
		// 				IntentContent: &armnetwork.IntentContent{
		// 					DestinationResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmDest"),
		// 					IPTraffic: &armnetwork.IPTraffic{
		// 						DestinationIPs: []*string{
		// 							to.Ptr("10.4.0.1")},
		// 							DestinationPorts: []*string{
		// 								to.Ptr("0")},
		// 								Protocols: []*armnetwork.NetworkProtocol{
		// 									to.Ptr(armnetwork.NetworkProtocolAny)},
		// 									SourceIPs: []*string{
		// 										to.Ptr("10.4.0.0")},
		// 										SourcePorts: []*string{
		// 											to.Ptr("0")},
		// 										},
		// 										SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmSrc"),
		// 									},
		// 									IntentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisIntents/testReachabilityAnalysisIntenant1"),
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								},
		// 						}},
		// 					}
	}
}
