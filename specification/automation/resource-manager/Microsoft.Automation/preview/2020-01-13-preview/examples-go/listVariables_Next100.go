package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listVariables_Next100.json
func ExampleVariableClient_NewListByAutomationAccountPager_listVariablesNext100() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVariableClient().NewListByAutomationAccountPager("rg", "sampleAccount9", nil)
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
		// page.VariableListResult = armautomation.VariableListResult{
		// 	Value: []*armautomation.Variable{
		// 		{
		// 			Name: to.Ptr("sampleVariable9"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable9"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:01:07.380Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:01:07.380Z"); return t}()),
		// 				Value: to.Ptr("\"server9.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable90"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable90"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:19.147Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:19.147Z"); return t}()),
		// 				Value: to.Ptr("\"server90.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable91"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable91"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:20.257Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:20.257Z"); return t}()),
		// 				Value: to.Ptr("\"server91.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable92"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable92"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:21.037Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:21.037Z"); return t}()),
		// 				Value: to.Ptr("\"server92.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable93"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable93"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:21.803Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:21.803Z"); return t}()),
		// 				Value: to.Ptr("\"server93.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable94"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable94"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:22.583Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:22.583Z"); return t}()),
		// 				Value: to.Ptr("\"server94.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable95"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable95"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:23.333Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:23.333Z"); return t}()),
		// 				Value: to.Ptr("\"server95.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable96"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable96"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:24.163Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:24.163Z"); return t}()),
		// 				Value: to.Ptr("\"server96.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable97"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable97"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:24.973Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:24.973Z"); return t}()),
		// 				Value: to.Ptr("\"server97.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable98"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable98"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:25.757Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:25.757Z"); return t}()),
		// 				Value: to.Ptr("\"server98.domain.com\""),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleVariable99"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable99"),
		// 			Properties: &armautomation.VariableProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:26.553Z"); return t}()),
		// 				IsEncrypted: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:02:26.553Z"); return t}()),
		// 				Value: to.Ptr("\"server99.domain.com\""),
		// 			},
		// 	}},
		// }
	}
}
