package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapDatabaseInstances_List.json
func ExampleSAPDatabaseInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("6d875e77-e412-4d7d-9af4-8895278b4443", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSAPDatabaseInstancesClient().NewListPager("test-rg", "X00", nil)
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
		// page = armworkloadssapvirtualinstance.SAPDatabaseInstancesClientListResponse{
		// 	SAPDatabaseInstanceListResult: armworkloadssapvirtualinstance.SAPDatabaseInstanceListResult{
		// 		Value: []*armworkloadssapvirtualinstance.SAPDatabaseInstance{
		// 			{
		// 				Name: to.Ptr("databaseServer"),
		// 				Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances"),
		// 				ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00/databaseInstances/databaseServer"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armworkloadssapvirtualinstance.SAPDatabaseProperties{
		// 					DatabaseSid: to.Ptr("X00"),
		// 					DatabaseType: to.Ptr("hdb"),
		// 					IPAddress: to.Ptr("10.0.0.5"),
		// 					ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
		// 					Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
		// 					Subnet: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
		// 					VMDetails: []*armworkloadssapvirtualinstance.DatabaseVMDetails{
		// 						{
		// 							VirtualMachineID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/db-vm"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armworkloadssapvirtualinstance.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@xyz.com"),
		// 					CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@xyz.com"),
		// 					LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
