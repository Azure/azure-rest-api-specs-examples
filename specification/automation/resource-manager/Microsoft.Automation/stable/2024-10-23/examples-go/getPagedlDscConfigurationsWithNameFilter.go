package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/getPagedlDscConfigurationsWithNameFilter.json
func ExampleDscConfigurationClient_NewListByAutomationAccountPager_listPagedDscConfigurationsWithNameFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscConfigurationClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscConfigurationClientListByAutomationAccountOptions{
		Filter:      to.Ptr("contains(name,'server')"),
		Inlinecount: to.Ptr("allpages"),
		Skip:        to.Ptr[int32](0),
		Top:         to.Ptr[int32](2)})
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
		// page = armautomation.DscConfigurationClientListByAutomationAccountResponse{
		// 	DscConfigurationListResult: armautomation.DscConfigurationListResult{
		// 		TotalCount: to.Ptr[int32](4),
		// 		Value: []*armautomation.DscConfiguration{
		// 			{
		// 				Name: to.Ptr("SqlServerBig"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Configurations"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/configurations/SqlServerBig"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.DscConfigurationProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T18:53:24.997+00:00"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T18:53:24.997+00:00"); return t}()),
		// 					NodeConfigurationCount: to.Ptr[int32](1),
		// 					State: to.Ptr(armautomation.DscConfigurationStatePublished),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SqlServerLittle"),
		// 				Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Configurations"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/configurations/SqlServerLittle"),
		// 				Location: to.Ptr("East US 2"),
		// 				Properties: &armautomation.DscConfigurationProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T18:53:24.997+00:00"); return t}()),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T18:53:24.997+00:00"); return t}()),
		// 					NodeConfigurationCount: to.Ptr[int32](1),
		// 					State: to.Ptr(armautomation.DscConfigurationStatePublished),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
