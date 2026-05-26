package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/updateAutomationAccount.json
func ExampleAccountClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountClient().Update(ctx, "rg", "myAutomationAccount9", armautomation.AccountUpdateParameters{
		Name:     to.Ptr("myAutomationAccount9"),
		Location: to.Ptr("East US 2"),
		Properties: &armautomation.AccountUpdateProperties{
			SKU: &armautomation.SKU{
				Name: to.Ptr(armautomation.SKUNameEnumFree),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.AccountClientUpdateResponse{
	// 	Account: armautomation.Account{
	// 		Name: to.Ptr("myAutomationAccount9"),
	// 		Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount9"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.AccountProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T01:13:43.267+00:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T01:13:43.267+00:00"); return t}()),
	// 			SKU: &armautomation.SKU{
	// 				Name: to.Ptr(armautomation.SKUNameEnumFree),
	// 			},
	// 			State: to.Ptr(armautomation.AutomationAccountStateOk),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
