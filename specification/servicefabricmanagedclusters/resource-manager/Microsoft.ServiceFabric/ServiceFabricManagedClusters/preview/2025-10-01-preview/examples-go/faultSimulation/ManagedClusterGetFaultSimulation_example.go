package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/faultSimulation/ManagedClusterGetFaultSimulation_example.json
func ExampleManagedClustersClient_GetFaultSimulation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetFaultSimulation(ctx, "resRg", "myCluster", armservicefabricmanagedclusters.FaultSimulationIDContent{
		SimulationID: to.Ptr("aec13cc2-1d39-4ba6-a1a8-2fc35b00643c"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ManagedClustersClientGetFaultSimulationResponse{
	// 	FaultSimulation: &armservicefabricmanagedclusters.FaultSimulation{
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00"); return t}()),
	// 		SimulationID: to.Ptr("aec13cc2-1d39-4ba6-a1a8-2fc35b00643c"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-21T18:11:11.3471742Z"); return t}()),
	// 		Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusActive),
	// 		Details: &armservicefabricmanagedclusters.FaultSimulationDetails{
	// 			OperationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ClusterID: to.Ptr("dd65fb6d-753b-4664-8798-4f077d4a2e18"),
	// 			NodeTypeFaultSimulation: []*armservicefabricmanagedclusters.NodeTypeFaultSimulation{
	// 				{
	// 					OperationID: to.Ptr("714f0003-80d8-464a-a019-69e6bf89ed4c"),
	// 					NodeTypeName: to.Ptr("BE"),
	// 					OperationStatus: to.Ptr(armservicefabricmanagedclusters.SfmcOperationStatusSucceeded),
	// 					Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusActive),
	// 				},
	// 			},
	// 			Parameters: &armservicefabricmanagedclusters.ZoneFaultSimulationContent{
	// 				Constraints: &armservicefabricmanagedclusters.FaultSimulationConstraints{
	// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00"); return t}()),
	// 				},
	// 				FaultKind: to.Ptr(armservicefabricmanagedclusters.FaultKindZone),
	// 				Force: to.Ptr(false),
	// 				Zones: []*string{
	// 					to.Ptr("2"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
