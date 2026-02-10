package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ReachabilityAnalysisRunPut.json
func ExampleReachabilityAnalysisRunsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReachabilityAnalysisRunsClient().Create(ctx, "rg1", "testNetworkManager", "testWorkspace", "testAnalysisRunName", armnetwork.ReachabilityAnalysisRun{
		Properties: &armnetwork.ReachabilityAnalysisRunProperties{
			Description: to.Ptr("A sample reachability analysis run"),
			IntentID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisIntents/testReachabilityAnalysisIntenant1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReachabilityAnalysisRun = armnetwork.ReachabilityAnalysisRun{
	// 	Name: to.Ptr("testAnalysisRunName"),
	// 	Type: to.Ptr("Microsoft.Network/networkManagers/verifierWorkspaces/reachabilityAnalysisRuns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testWorkspace/reachabilityAnalysisRuns/testAnalysisRunName"),
	// 	SystemData: &armnetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
	// 		LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnetwork.ReachabilityAnalysisRunProperties{
	// 		Description: to.Ptr("A sample  reachability analysis intent"),
	// 		IntentContent: &armnetwork.IntentContent{
	// 			DestinationResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmDest"),
	// 			IPTraffic: &armnetwork.IPTraffic{
	// 				DestinationIPs: []*string{
	// 					to.Ptr("10.4.0.1")},
	// 					DestinationPorts: []*string{
	// 						to.Ptr("0")},
	// 						Protocols: []*armnetwork.NetworkProtocol{
	// 							to.Ptr(armnetwork.NetworkProtocolAny)},
	// 							SourceIPs: []*string{
	// 								to.Ptr("10.4.0.0")},
	// 								SourcePorts: []*string{
	// 									to.Ptr("0")},
	// 								},
	// 								SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/testVmSrc"),
	// 							},
	// 							IntentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/testNetworkManager/verifierWorkspaces/testVerifierWorkspace1/reachabilityAnalysisIntents/testReachabilityAnalysisIntenant1"),
	// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						},
	// 					}
}
