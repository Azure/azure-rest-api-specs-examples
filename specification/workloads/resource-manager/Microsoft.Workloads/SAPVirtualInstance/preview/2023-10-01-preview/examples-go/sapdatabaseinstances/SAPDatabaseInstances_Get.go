package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapdatabaseinstances/SAPDatabaseInstances_Get.json
func ExampleSAPDatabaseInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.SAPDatabaseInstance = armworkloadssapvirtualinstance.SAPDatabaseInstance{
	// 	Name: to.Ptr("databaseServer"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/databaseInstances/databaseServer"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPDatabaseProperties{
	// 		DatabaseSid: to.Ptr("X00"),
	// 		DatabaseType: to.Ptr("hdb"),
	// 		IPAddress: to.Ptr("10.0.0.5"),
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
	// 		Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		VMDetails: []*armworkloadssapvirtualinstance.DatabaseVMDetails{
	// 			{
	// 				VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/db-vm"),
	// 		}},
	// 	},
	// }
}
