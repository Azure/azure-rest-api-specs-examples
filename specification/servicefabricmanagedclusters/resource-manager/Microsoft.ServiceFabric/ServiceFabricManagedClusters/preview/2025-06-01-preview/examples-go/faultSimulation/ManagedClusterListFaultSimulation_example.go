package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-06-01-preview/faultSimulation/ManagedClusterListFaultSimulation_example.json
func ExampleManagedClustersClient_NewListFaultSimulationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListFaultSimulationPager("resRg", "myCluster", nil)
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
		// page = armservicefabricmanagedclusters.ManagedClustersClientListFaultSimulationResponse{
		// 	FaultSimulationListResult: armservicefabricmanagedclusters.FaultSimulationListResult{
		// 		Value: []*armservicefabricmanagedclusters.FaultSimulation{
		// 			{
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00"); return t}()),
		// 				SimulationID: to.Ptr("1bb61ba9-8a41-4d73-b5f0-7fc93b1edfe3"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-21T21:27:55.4452675Z"); return t}()),
		// 				Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusActive),
		// 				Details: &armservicefabricmanagedclusters.FaultSimulationDetails{
		// 					OperationID: to.Ptr("b7997d2f-1f07-4245-b197-137b196ecaf3"),
		// 					ClusterID: to.Ptr("dd65fb6d-753b-4664-8798-4f077d4a2e18"),
		// 					NodeTypeFaultSimulation: []*armservicefabricmanagedclusters.NodeTypeFaultSimulation{
		// 						{
		// 							OperationID: to.Ptr("bff10003-af96-465c-b849-99c18e1f4af6"),
		// 							NodeTypeName: to.Ptr("BE"),
		// 							OperationStatus: to.Ptr(armservicefabricmanagedclusters.SfmcOperationStatusSucceeded),
		// 							Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusActive),
		// 						},
		// 						{
		// 							OperationID: to.Ptr("6021d3c9-6def-4e9e-b133-d58ac1c3a4cc"),
		// 							NodeTypeName: to.Ptr("FE"),
		// 							OperationStatus: to.Ptr(armservicefabricmanagedclusters.SfmcOperationStatusSucceeded),
		// 							Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusActive),
		// 						},
		// 					},
		// 					Parameters: &armservicefabricmanagedclusters.ZoneFaultSimulationContent{
		// 						Constraints: &armservicefabricmanagedclusters.FaultSimulationConstraints{
		// 							ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00"); return t}()),
		// 						},
		// 						FaultKind: to.Ptr(armservicefabricmanagedclusters.FaultKindZone),
		// 						Force: to.Ptr(false),
		// 						Zones: []*string{
		// 							to.Ptr("3"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-21T21:24:13.082339Z"); return t}()),
		// 				SimulationID: to.Ptr("aec13cc2-1d39-4ba6-a1a8-2fc35b00643c"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-21T18:11:11.3471742Z"); return t}()),
		// 				Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusDone),
		// 				Details: &armservicefabricmanagedclusters.FaultSimulationDetails{
		// 					OperationID: to.Ptr("3e22a9a7-13c1-450b-af5e-2c739b963bdf"),
		// 					ClusterID: to.Ptr("dd65fb6d-753b-4664-8798-4f077d4a2e18"),
		// 					NodeTypeFaultSimulation: []*armservicefabricmanagedclusters.NodeTypeFaultSimulation{
		// 						{
		// 							OperationID: to.Ptr("391d0003-0d46-474c-9839-cbc345938704"),
		// 							NodeTypeName: to.Ptr("BE"),
		// 							OperationStatus: to.Ptr(armservicefabricmanagedclusters.SfmcOperationStatusSucceeded),
		// 							Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusDone),
		// 						},
		// 						{
		// 							OperationID: to.Ptr("cf4ac1d7-e588-491e-9e6f-ce15a17bf389"),
		// 							NodeTypeName: to.Ptr("FE"),
		// 							OperationStatus: to.Ptr(armservicefabricmanagedclusters.SfmcOperationStatusSucceeded),
		// 							Status: to.Ptr(armservicefabricmanagedclusters.FaultSimulationStatusDone),
		// 						},
		// 					},
		// 					Parameters: &armservicefabricmanagedclusters.ZoneFaultSimulationContent{
		// 						Constraints: &armservicefabricmanagedclusters.FaultSimulationConstraints{
		// 							ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00"); return t}()),
		// 						},
		// 						FaultKind: to.Ptr(armservicefabricmanagedclusters.FaultKindZone),
		// 						Force: to.Ptr(false),
		// 						Zones: []*string{
		// 							to.Ptr("2"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
