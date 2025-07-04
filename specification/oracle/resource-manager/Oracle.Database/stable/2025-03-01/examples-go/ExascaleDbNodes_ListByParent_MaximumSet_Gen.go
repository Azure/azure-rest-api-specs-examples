package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/ExascaleDbNodes_ListByParent_MaximumSet_Gen.json
func ExampleExascaleDbNodesClient_NewListByParentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExascaleDbNodesClient().NewListByParentPager("rgopenapi", "vmClusterName", nil)
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
		// page = armoracledatabase.ExascaleDbNodesClientListByParentResponse{
		// 	ExascaleDbNodeListResult: armoracledatabase.ExascaleDbNodeListResult{
		// 		Value: []*armoracledatabase.ExascaleDbNode{
		// 			{
		// 				Properties: &armoracledatabase.ExascaleDbNodeProperties{
		// 					Ocid: to.Ptr("ocid1.dbnode.oc1..aaaaa3klq"),
		// 					AdditionalDetails: to.Ptr("zjvaydzrzxrmtiolutkhyfumql"),
		// 					CPUCoreCount: to.Ptr[int32](25),
		// 					DbNodeStorageSizeInGbs: to.Ptr[int32](7),
		// 					FaultDomain: to.Ptr("bgtzblfwbdooaj"),
		// 					Hostname: to.Ptr("nmbmxqpkdqueswkwystaupanqrn"),
		// 					LifecycleState: to.Ptr(armoracledatabase.DbNodeProvisioningStateAvailable),
		// 					MaintenanceType: to.Ptr("ncsgznwyxmzcrqnmzbn"),
		// 					MemorySizeInGbs: to.Ptr[int32](29),
		// 					SoftwareStorageSizeInGb: to.Ptr[int32](14),
		// 					TimeMaintenanceWindowEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:38.078Z"); return t}()),
		// 					TimeMaintenanceWindowStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:38.078Z"); return t}()),
		// 					TotalCPUCoreCount: to.Ptr[int32](26),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Oracle.Database/exadbVmClusters/vmCluster/dbNodes/dbNodeName"),
		// 				Name: to.Ptr("lkjpzwgzy"),
		// 				Type: to.Ptr("zdrljrxhtseejhwvzox"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ilrpjodjmvzhybazxipoplnql"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lhjbxchqkaia"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-09T21:02:12.592Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ajhmfpn"),
		// 	},
		// }
	}
}
