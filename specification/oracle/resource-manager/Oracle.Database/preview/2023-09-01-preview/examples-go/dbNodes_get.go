package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbNodes_get.json
func ExampleDbNodesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDbNodesClient().Get(ctx, "rg000", "cluster1", "ocid1....aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DbNode = armoracledatabase.DbNode{
	// 	Type: to.Ptr("Oracle.Database/cloudVmClusters/dbNodes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1/dbNodes/ocid1"),
	// 	Properties: &armoracledatabase.DbNodeProperties{
	// 		BackupIPID: to.Ptr("id1"),
	// 		BackupVnic2ID: to.Ptr("id1"),
	// 		BackupVnicID: to.Ptr("id1"),
	// 		CPUCoreCount: to.Ptr[int32](1000),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](500),
	// 		DbServerID: to.Ptr("dbserver1"),
	// 		DbSystemID: to.Ptr("db1"),
	// 		FaultDomain: to.Ptr("domain1"),
	// 		HostIPID: to.Ptr("10.0.0.0"),
	// 		Hostname: to.Ptr("host1"),
	// 		LifecycleState: to.Ptr(armoracledatabase.DbNodeProvisioningStateAvailable),
	// 		MaintenanceType: to.Ptr(armoracledatabase.DbNodeMaintenanceTypeVmdbRebootMigration),
	// 		MemorySizeInGbs: to.Ptr[int32](100),
	// 		Ocid: to.Ptr("ocid.dbNodes.1"),
	// 		SoftwareStorageSizeInGb: to.Ptr[int32](1000),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 		TimeMaintenanceWindowEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 		TimeMaintenanceWindowStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 		Vnic2ID: to.Ptr("ocid.vnic.id2"),
	// 		VnicID: to.Ptr("ocid.vnic.id1"),
	// 	},
	// }
}
