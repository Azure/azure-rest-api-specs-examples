package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
)

// Generated from example definition: 2025-09-01/DbNodes_Action_MaximumSet_Gen.json
func ExampleDbNodesClient_BeginAction_vmActionsOnDbNodesOfVMClusterGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDbNodesClient().BeginAction(ctx, "rgopenapi", "cloudvmcluster1", "abciderewdidsereq", armoracledatabase.DbNodeAction{
		Action: to.Ptr(armoracledatabase.DbNodeActionEnumStart),
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
	// res = armoracledatabase.DbNodesClientActionResponse{
	// 	DbNode: &armoracledatabase.DbNode{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1/dbNodes/ocid1"),
	// 		Type: to.Ptr("Oracle.Database/cloudVmClusters/dbNodes"),
	// 		Properties: &armoracledatabase.DbNodeProperties{
	// 			Ocid: to.Ptr("ocid.dbNodes.1"),
	// 			BackupIPID: to.Ptr("id1"),
	// 			BackupVnic2ID: to.Ptr("id1"),
	// 			BackupVnicID: to.Ptr("id1"),
	// 			CPUCoreCount: to.Ptr[int32](1000),
	// 			DbNodeStorageSizeInGbs: to.Ptr[int32](500),
	// 			DbServerID: to.Ptr("dbserver1"),
	// 			DbSystemID: to.Ptr("db1"),
	// 			FaultDomain: to.Ptr("domain1"),
	// 			HostIPID: to.Ptr("10.0.0.0"),
	// 			Hostname: to.Ptr("host1"),
	// 			LifecycleState: to.Ptr(armoracledatabase.DbNodeProvisioningStateAvailable),
	// 			MaintenanceType: to.Ptr(armoracledatabase.DbNodeMaintenanceTypeVmdbRebootMigration),
	// 			MemorySizeInGbs: to.Ptr[int32](100),
	// 			SoftwareStorageSizeInGb: to.Ptr[int32](1000),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 			TimeMaintenanceWindowEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 			TimeMaintenanceWindowStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-21T13:44:04.924Z"); return t}()),
	// 			Vnic2ID: to.Ptr("ocid.vnic.id2"),
	// 			VnicID: to.Ptr("ocid.vnic.id1"),
	// 			AdditionalDetails: to.Ptr("ympoempzdwoemuoosvbznsazqxruk"),
	// 			LifecycleDetails: to.Ptr("ztozxox"),
	// 			ProvisioningState: to.Ptr(armoracledatabase.ResourceProvisioningStateSucceeded),
	// 		},
	// 		Name: to.Ptr("fwzoiphvtqxgvyvdkzdjhjtemeifj"),
	// 		SystemData: &armoracledatabase.SystemData{
	// 			CreatedBy: to.Ptr("sqehacivpuim"),
	// 			CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("axrqfdkqylvjv"),
	// 			LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-01T04:32:58.716Z"); return t}()),
	// 		},
	// 	},
	// }
}
