package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithCompositeFilter.json
func ExampleDscNodeClient_NewListByAutomationAccountPager_listPagedDscNodesWithFiltersSeparatedByAnd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscNodeClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscNodeClientListByAutomationAccountOptions{Filter: to.Ptr("properties/extensionHandler/any(eh: eh/version gt '2.70') and contains(name,'sql') and contains(properties/nodeConfiguration/name,'$$Not$$Configured$$')"),
		Skip:        to.Ptr[int32](0),
		Top:         to.Ptr[int32](10),
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
		// page.DscNodeListResult = armautomation.DscNodeListResult{
		// 	TotalCount: to.Ptr[int32](1),
		// 	Value: []*armautomation.DscNode{
		// 		{
		// 			Name: to.Ptr("Sql1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Nodes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/Sql1"),
		// 			Properties: &armautomation.DscNodeProperties{
		// 				ExtensionHandler: []*armautomation.DscNodeExtensionHandlerAssociationProperty{
		// 					{
		// 						Name: to.Ptr("Microsoft.Powershell.DSC"),
		// 						Version: to.Ptr("2.70.0.1"),
		// 				}},
		// 				IP: to.Ptr("ip"),
		// 				LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-22T22:25:39.096Z"); return t}()),
		// 				NodeConfiguration: &armautomation.DscNodeConfigurationAssociationProperty{
		// 					Name: to.Ptr("SetupSqlServer.localhost"),
		// 				},
		// 				NodeID: to.Ptr("01D64CDE-85DD-4C9B-B8F1-2F725348FDEC"),
		// 				RegistrationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T00:51:12.539Z"); return t}()),
		// 				Status: to.Ptr("Compliant"),
		// 			},
		// 	}},
		// }
	}
}
