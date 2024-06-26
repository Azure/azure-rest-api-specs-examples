package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Monitor/preview/2021-06-03-preview/examples/OperationsGet.json
func ExampleOperationsForMonitorClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsForMonitorClient().NewListPager(nil)
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
		// page.OperationListResultAutoGenerated = armmonitor.OperationListResultAutoGenerated{
		// 	Value: []*armmonitor.OperationAutoGenerated{
		// 		{
		// 			Name: to.Ptr("microsoft.monitor/accounts/read"),
		// 			Display: &armmonitor.OperationDisplayAutoGenerated{
		// 				Description: to.Ptr("Read any Azure Monitor Workspace"),
		// 				Operation: to.Ptr("Read Azure Monitor Workspaces"),
		// 				Provider: to.Ptr("Microsoft Monitoring"),
		// 				Resource: to.Ptr("Azure Monitor Workspaces"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armmonitor.Origin("user, system")),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.monitor/accounts/write"),
		// 			Display: &armmonitor.OperationDisplayAutoGenerated{
		// 				Description: to.Ptr("Create or Update any Azure Monitor Workspace"),
		// 				Operation: to.Ptr("Create or Update Azure Monitor Workspaces"),
		// 				Provider: to.Ptr("Microsoft Monitoring"),
		// 				Resource: to.Ptr("Azure Monitor Workspaces"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armmonitor.Origin("user, system")),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.monitor/accounts/delete"),
		// 			Display: &armmonitor.OperationDisplayAutoGenerated{
		// 				Description: to.Ptr("Delete any Azure Monitor Workspace"),
		// 				Operation: to.Ptr("Delete Azure Monitor Workspaces"),
		// 				Provider: to.Ptr("Microsoft Monitoring"),
		// 				Resource: to.Ptr("Azure Monitor Workspaces"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armmonitor.Origin("user, system")),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.monitor/accounts/metrics/read"),
		// 			Display: &armmonitor.OperationDisplayAutoGenerated{
		// 				Description: to.Ptr("Read Azure Monitor Workspace metrics"),
		// 				Operation: to.Ptr("Read Azure Monitor Workspace metrics"),
		// 				Provider: to.Ptr("Microsoft Monitoring"),
		// 				Resource: to.Ptr("Azure Monitor Workspaces"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armmonitor.Origin("user, system")),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.monitor/accounts/data/metrics/read"),
		// 			Display: &armmonitor.OperationDisplayAutoGenerated{
		// 				Description: to.Ptr("Read metrics data in any Azure Monitor Workspace"),
		// 				Operation: to.Ptr("Read Metrics Data"),
		// 				Provider: to.Ptr("Microsoft Monitoring"),
		// 				Resource: to.Ptr("Azure Monitor Workspaces"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr(armmonitor.Origin("user, system")),
		// 	}},
		// }
	}
}
