package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listDscNodeConfigurations.json
func ExampleDscNodeConfigurationClient_NewListByAutomationAccountPager_listDscNodeConfigurationsByAutomationAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscNodeConfigurationClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", nil)
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
		// page = armautomation.DscNodeConfigurationClientListByAutomationAccountResponse{
		// 	DscNodeConfigurationListResult: armautomation.DscNodeConfigurationListResult{
		// 		Value: []*armautomation.DscNodeConfiguration{
		// 			{
		// 				Name: to.Ptr("SetupServer.localhost"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/nodeConfigurations"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodeConfigurations/SetupServer.localhost"),
		// 				Properties: &armautomation.DscNodeConfigurationProperties{
		// 					Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 						Name: to.Ptr("SetupServer"),
		// 					},
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
		// 					IncrementNodeConfigurationBuild: to.Ptr(false),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
		// 					NodeCount: to.Ptr[int64](0),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SetupServer.localhost"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/nodeConfigurations"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodeConfigurations/SetupServer.localhost"),
		// 				Properties: &armautomation.DscNodeConfigurationProperties{
		// 					Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 						Name: to.Ptr("SetupServer"),
		// 					},
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
		// 					IncrementNodeConfigurationBuild: to.Ptr(false),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.8901968+00:00"); return t}()),
		// 					NodeCount: to.Ptr[int64](0),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
