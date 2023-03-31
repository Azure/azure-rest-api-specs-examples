package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_List.json
func ExampleSapMonitorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhanaonazure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSapMonitorsClient().NewListPager(nil)
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
		// page.SapMonitorListResult = armhanaonazure.SapMonitorListResult{
		// 	Value: []*armhanaonazure.SapMonitor{
		// 		{
		// 			Name: to.Ptr("mySapMonitor1"),
		// 			Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armhanaonazure.SapMonitorProperties{
		// 				EnableCustomerAnalytics: to.Ptr(true),
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				LogAnalyticsWorkspaceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				LogAnalyticsWorkspaceSharedKey: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000=="),
		// 				ManagedResourceGroupName: to.Ptr("myManagedResourceGroup"),
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				ProvisioningState: to.Ptr(armhanaonazure.HanaProvisioningStatesEnumSucceeded),
		// 				SapMonitorCollectorVersion: to.Ptr("v1.0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySapMonitor2"),
		// 			Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armhanaonazure.SapMonitorProperties{
		// 				EnableCustomerAnalytics: to.Ptr(true),
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				LogAnalyticsWorkspaceID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				LogAnalyticsWorkspaceSharedKey: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000=="),
		// 				ManagedResourceGroupName: to.Ptr("myManagedResourceGroup"),
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				ProvisioningState: to.Ptr(armhanaonazure.HanaProvisioningStatesEnumSucceeded),
		// 				SapMonitorCollectorVersion: to.Ptr("v1.0"),
		// 			},
		// 	}},
		// }
	}
}
