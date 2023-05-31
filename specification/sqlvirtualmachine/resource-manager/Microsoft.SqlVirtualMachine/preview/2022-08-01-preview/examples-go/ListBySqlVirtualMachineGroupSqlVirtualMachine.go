package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79a5aa63c0551c1b5af1d2853cceb495283d334/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/ListBySqlVirtualMachineGroupSqlVirtualMachine.json
func ExampleSQLVirtualMachinesClient_NewListBySQLVMGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLVirtualMachinesClient().NewListBySQLVMGroupPager("testrg", "testvm", nil)
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
		// page.ListResult = armsqlvirtualmachine.ListResult{
		// 	Value: []*armsqlvirtualmachine.SQLVirtualMachine{
		// 		{
		// 			Name: to.Ptr("testvm2"),
		// 			Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachines"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
		// 			Location: to.Ptr("northeurope"),
		// 			Properties: &armsqlvirtualmachine.Properties{
		// 				EnableAutomaticUpgrade: to.Ptr(false),
		// 				LeastPrivilegeMode: to.Ptr(armsqlvirtualmachine.LeastPrivilegeModeNotSet),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLImageOffer: to.Ptr("SQL2014-WS2012R2"),
		// 				SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLImageSKUEnterprise),
		// 				SQLManagement: to.Ptr(armsqlvirtualmachine.SQLManagementModeFull),
		// 				SQLServerLicenseType: to.Ptr(armsqlvirtualmachine.SQLServerLicenseTypePAYG),
		// 				SQLVirtualMachineGroupResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup"),
		// 				VirtualMachineResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm2"),
		// 			},
		// 	}},
		// }
	}
}
