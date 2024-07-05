package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbServers_get.json
func ExampleDbServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDbServersClient().Get(ctx, "rg000", "infra1", "ocid1....aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DbServer = armoracledatabase.DbServer{
	// 	Type: to.Ptr("Oracle.Database/cloudVmClusters/dbServers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1/dbServers/ocid1"),
	// 	Properties: &armoracledatabase.DbServerProperties{
	// 		AutonomousVirtualMachineIDs: []*string{
	// 			to.Ptr("ocid1..aaaaa")},
	// 			AutonomousVMClusterIDs: []*string{
	// 				to.Ptr("ocid1..aaaaa")},
	// 				CompartmentID: to.Ptr("ocid1....aaaa"),
	// 				CPUCoreCount: to.Ptr[int32](100),
	// 				DbNodeIDs: []*string{
	// 					to.Ptr("ocid1..aaaaa")},
	// 					DbNodeStorageSizeInGbs: to.Ptr[int32](150),
	// 					DisplayName: to.Ptr("dbserver1"),
	// 					ExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 					LifecycleState: to.Ptr(armoracledatabase.DbServerProvisioningStateAvailable),
	// 					MaxCPUCount: to.Ptr[int32](1000),
	// 					MaxMemoryInGbs: to.Ptr[int32](1000),
	// 					Ocid: to.Ptr("ocid1"),
	// 					VMClusterIDs: []*string{
	// 						to.Ptr("ocid1..aaaaa")},
	// 					},
	// 				}
}
