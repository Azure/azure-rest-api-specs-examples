package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeConfigurationsWithNameFilter.json
func ExampleDscNodeConfigurationClient_NewListByAutomationAccountPager_listPagedDscNodeConfigurationsByAutomationAccountWithNameFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscNodeConfigurationClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscNodeConfigurationClientListByAutomationAccountOptions{Filter: to.Ptr("contains('.localhost',name)"),
		Skip:        to.Ptr[int32](0),
		Top:         to.Ptr[int32](2),
		Inlinecount: to.Ptr("allpages"),
	})
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
		// page.DscNodeConfigurationListResult = armautomation.DscNodeConfigurationListResult{
		// 	TotalCount: to.Ptr[int32](6),
		// 	Value: []*armautomation.DscNodeConfiguration{
		// 		{
		// 			Name: to.Ptr("server.localhost"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/nodeConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodeConfigurations/SetupServer.localhost"),
		// 			Properties: &armautomation.DscNodeConfigurationProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("SetupServer"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.890Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.890Z"); return t}()),
		// 				NodeCount: to.Ptr[int64](2),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SetupClient.localhost"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/nodeConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodeConfigurations/SetupClient.localhost"),
		// 			Properties: &armautomation.DscNodeConfigurationProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("SetupClient"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.890Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:17:06.890Z"); return t}()),
		// 				NodeCount: to.Ptr[int64](6),
		// 			},
		// 	}},
		// }
	}
}
