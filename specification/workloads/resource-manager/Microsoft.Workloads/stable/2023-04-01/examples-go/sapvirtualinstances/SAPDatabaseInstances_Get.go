package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPDatabaseInstances_Get.json
func ExampleSAPDatabaseInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPDatabaseInstancesClient().Get(ctx, "test-rg", "X00", "databaseServer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPDatabaseInstance = armworkloads.SAPDatabaseInstance{
	// 	Name: to.Ptr("databaseServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/databaseInstances/databaseServer"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloads.SAPDatabaseProperties{
	// 		DatabaseSid: to.Ptr("X00"),
	// 		DatabaseType: to.Ptr("hdb"),
	// 		IPAddress: to.Ptr("10.0.0.5"),
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		VMDetails: []*armworkloads.DatabaseVMDetails{
	// 			{
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/db-vm"),
	// 		}},
	// 	},
	// }
}
