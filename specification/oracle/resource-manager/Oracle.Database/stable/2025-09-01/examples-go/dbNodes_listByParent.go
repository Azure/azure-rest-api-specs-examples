package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/dbNodes_listByParent.json
func ExampleDbNodesClient_NewListByCloudVMClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbNodesClient().NewListByCloudVMClusterPager("rg000", "cluster1", nil)
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
		// page = armoracledatabase.DbNodesClientListByCloudVMClusterResponse{
		// 	DbNodeListResult: armoracledatabase.DbNodeListResult{
		// 		Value: []*armoracledatabase.DbNode{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1/dbNodes/ocid1"),
		// 				Type: to.Ptr("Oracle.Database/cloudVmClusters/dbNodes"),
		// 				Properties: &armoracledatabase.DbNodeProperties{
		// 					Ocid: to.Ptr("ocid.dbNodes.1"),
		// 					BackupIPID: to.Ptr("id1"),
		// 					BackupVnic2ID: to.Ptr("id1"),
		// 					BackupVnicID: to.Ptr("id1"),
		// 					CPUCoreCount: to.Ptr[int32](1000),
		// 					DbNodeStorageSizeInGbs: to.Ptr[int32](500),
		// 					DbServerID: to.Ptr("dbserver1"),
		// 					DbSystemID: to.Ptr("db1"),
		// 					FaultDomain: to.Ptr("domain1"),
		// 					HostIPID: to.Ptr("10.0.0.0"),
		// 					Hostname: to.Ptr("host1"),
		// 					LifecycleState: to.Ptr(armoracledatabase.DbNodeProvisioningStateAvailable),
		// 					MaintenanceType: to.Ptr(armoracledatabase.DbNodeMaintenanceTypeVmdbRebootMigration),
		// 					MemorySizeInGbs: to.Ptr[int32](100),
		// 					SoftwareStorageSizeInGb: to.Ptr[int32](1000),
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
		// 					TimeMaintenanceWindowEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
		// 					TimeMaintenanceWindowStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
		// 					Vnic2ID: to.Ptr("ocid.vnic.id2"),
		// 					VnicID: to.Ptr("ocid.vnic.id1"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
