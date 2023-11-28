package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodesByAutomationAccountWithNameFilter.json
func ExampleDscNodeClient_NewListByAutomationAccountPager_listPagedDscNodesByAutomationAccountWithNameFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscNodeClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscNodeClientListByAutomationAccountOptions{Filter: to.Ptr("contains('DSCCOMP',name)"),
		Skip:        to.Ptr[int32](0),
		Top:         to.Ptr[int32](6),
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
		// 	TotalCount: to.Ptr[int32](2),
		// 	Value: []*armautomation.DscNode{
		// 		{
		// 			Name: to.Ptr("DSCCOMP"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Nodes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/DSCCOMP"),
		// 			Properties: &armautomation.DscNodeProperties{
		// 				ExtensionHandler: []*armautomation.DscNodeExtensionHandlerAssociationProperty{
		// 					{
		// 						Name: to.Ptr("Microsoft.Powershell.DSC"),
		// 						Version: to.Ptr("2.75.0.0"),
		// 				}},
		// 				IP: to.Ptr("ip"),
		// 				LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-22T22:25:39.096Z"); return t}()),
		// 				NodeConfiguration: &armautomation.DscNodeConfigurationAssociationProperty{
		// 					Name: to.Ptr("SetupServer.localhost"),
		// 				},
		// 				NodeID: to.Ptr("708D250A-2169-4B54-89FF-76F5F71C252A"),
		// 				RegistrationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T00:51:12.539Z"); return t}()),
		// 				Status: to.Ptr("Compliant"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DSCCOMP2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Nodes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/DSCCOMP2"),
		// 			Properties: &armautomation.DscNodeProperties{
		// 				ExtensionHandler: []*armautomation.DscNodeExtensionHandlerAssociationProperty{
		// 					{
		// 						Name: to.Ptr("Microsoft.Powershell.DSC"),
		// 						Version: to.Ptr("2.75.0.0"),
		// 				}},
		// 				IP: to.Ptr("ip"),
		// 				LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-22T22:25:39.096Z"); return t}()),
		// 				NodeConfiguration: &armautomation.DscNodeConfigurationAssociationProperty{
		// 					Name: to.Ptr("SetupServer.localhost"),
		// 				},
		// 				NodeID: to.Ptr("3B4BB31A-5132-4669-A15F-A17E234D1634"),
		// 				RegistrationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T00:51:12.539Z"); return t}()),
		// 				Status: to.Ptr("Failed"),
		// 			},
		// 	}},
		// }
	}
}
