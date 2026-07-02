package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: 2023-10-01/UpdateSqlVirtualMachineGroup.json
func ExampleGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGroupsClient().BeginUpdate(ctx, "testrg", "testvmgroup", armsqlvirtualmachine.GroupUpdate{
		Tags: map[string]*string{
			"mytag": to.Ptr("myval"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsqlvirtualmachine.GroupsClientUpdateResponse{
	// 	Group: armsqlvirtualmachine.Group{
	// 		Name: to.Ptr("testvmgroup"),
	// 		Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm"),
	// 		Location: to.Ptr("northeurope"),
	// 		Properties: &armsqlvirtualmachine.GroupProperties{
	// 			ProvisioningState: to.Ptr("UpdatingDomainful"),
	// 			SQLImageOffer: to.Ptr("SQL2017-WS2016"),
	// 			SQLImageSKU: to.Ptr(armsqlvirtualmachine.SQLVMGroupImageSKUEnterprise),
	// 			WsfcDomainProfile: &armsqlvirtualmachine.WsfcDomainProfile{
	// 				StorageAccountURL: to.Ptr("https://storgact.blob.core.windows.net/"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"mytag": to.Ptr("myval"),
	// 		},
	// 	},
	// }
}
