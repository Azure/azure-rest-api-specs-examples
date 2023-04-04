package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateDscNode.json
func ExampleDscNodeClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscNodeClient().Update(ctx, "rg", "myAutomationAccount33", "nodeId", armautomation.DscNodeUpdateParameters{
		NodeID: to.Ptr("nodeId"),
		Properties: &armautomation.DscNodeUpdateParametersProperties{
			NodeConfiguration: &armautomation.DscNodeConfigurationAssociationProperty{
				Name: to.Ptr("SetupServer.localhost"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscNode = armautomation.DscNode{
	// 	Name: to.Ptr("DSCCOMP"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Nodes"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId"),
	// 	Properties: &armautomation.DscNodeProperties{
	// 		ExtensionHandler: []*armautomation.DscNodeExtensionHandlerAssociationProperty{
	// 			{
	// 				Name: to.Ptr("Microsoft.Powershell.DSC"),
	// 				Version: to.Ptr("2.75.0.0"),
	// 		}},
	// 		IP: to.Ptr("ip"),
	// 		LastSeen: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-22T22:25:39.0963773+00:00"); return t}()),
	// 		NodeConfiguration: &armautomation.DscNodeConfigurationAssociationProperty{
	// 			Name: to.Ptr("SetupServer.localhost"),
	// 		},
	// 		NodeID: to.Ptr("nodeId"),
	// 		RegistrationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-10T00:51:12.5393083+00:00"); return t}()),
	// 		Status: to.Ptr("Pending"),
	// 	},
	// }
}
