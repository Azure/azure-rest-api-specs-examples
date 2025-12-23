package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/LinkedServicesListByWorkspace.json
func ExampleLinkedServicesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkedServicesClient().NewListByWorkspacePager("mms-eus", "TestLinkWS", nil)
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
		// page.LinkedServiceListResult = armoperationalinsights.LinkedServiceListResult{
		// 	Value: []*armoperationalinsights.LinkedService{
		// 		{
		// 			Name: to.Ptr("TestLinkWS/Automation"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testlinkws/linkedservices/automation"),
		// 			Properties: &armoperationalinsights.LinkedServiceProperties{
		// 				ProvisioningState: to.Ptr(armoperationalinsights.LinkedServiceEntityStatusSucceeded),
		// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/TestAccount"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestLinkWS/Cluster"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testlinkws/linkedservices/cluster"),
		// 			Properties: &armoperationalinsights.LinkedServiceProperties{
		// 				ProvisioningState: to.Ptr(armoperationalinsights.LinkedServiceEntityStatusSucceeded),
		// 				WriteAccessResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourceGroups/mms-eus/providers/Microsoft.OperationalInsights/clusters/TestCluster"),
		// 			},
		// 	}},
		// }
	}
}
