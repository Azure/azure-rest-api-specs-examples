package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listAutomationAccountKeys.json
func ExampleKeysClient_ListByAutomationAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKeysClient().ListByAutomationAccount(ctx, "rg", "MyAutomationAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.KeysClientListByAutomationAccountResponse{
	// 	KeyListResult: armautomation.KeyListResult{
	// 		Keys: []*armautomation.Key{
	// 			{
	// 				KeyName: to.Ptr(armautomation.AutomationKeyNamePrimary),
	// 				Permissions: to.Ptr(armautomation.AutomationKeyPermissionsFull),
	// 				Value: to.Ptr("**************************************************************"),
	// 			},
	// 			{
	// 				KeyName: to.Ptr(armautomation.AutomationKeyNameSecondary),
	// 				Permissions: to.Ptr(armautomation.AutomationKeyPermissionsFull),
	// 				Value: to.Ptr("**************************************************************"),
	// 			},
	// 		},
	// 	},
	// }
}
