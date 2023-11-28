package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/updateAutomationAccount.json
func ExampleAccountClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.Account = armautomation.Account{
	// 	Name: to.Ptr("myAutomationAccount9"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount9"),
	// 	Location: to.Ptr("East US 2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armautomation.AccountProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T01:13:43.267Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("myEmailId@microsoft.com"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T01:13:43.267Z"); return t}()),
	// 		SKU: &armautomation.SKU{
	// 			Name: to.Ptr(armautomation.SKUNameEnumFree),
	// 		},
	// 		State: to.Ptr(armautomation.AutomationAccountStateOk),
	// 	},
	// }
}
