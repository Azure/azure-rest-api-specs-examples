package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_List.json
func ExampleMonitorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListPager(nil)
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
		// page.MonitorListResult = armworkloads.MonitorListResult{
		// 	Value: []*armworkloads.Monitor{
		// 		{
		// 			Name: to.Ptr("mySapMonitor1"),
		// 			Type: to.Ptr("Microsoft.Workloads/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armworkloads.MonitorProperties{
		// 				AppLocation: to.Ptr("eastus"),
		// 				Errors: &armworkloads.MonitorPropertiesErrors{
		// 				},
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
		// 					Name: to.Ptr("myManagedRg1"),
		// 				},
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				MsiArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myMsi"),
		// 				ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				RoutingPreference: to.Ptr(armworkloads.RoutingPreferenceRouteAll),
		// 				StorageAccountArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		// 				ZoneRedundancyPreference: to.Ptr("ZoneRedundantApp"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySapMonitor2"),
		// 			Type: to.Ptr("Microsoft.Workloads/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armworkloads.MonitorProperties{
		// 				AppLocation: to.Ptr("westus"),
		// 				Errors: &armworkloads.MonitorPropertiesErrors{
		// 				},
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
		// 					Name: to.Ptr("myManagedRg2"),
		// 				},
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				MsiArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myMsi"),
		// 				ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				RoutingPreference: to.Ptr(armworkloads.RoutingPreferenceDefault),
		// 				StorageAccountArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		// 			},
		// 	}},
		// }
	}
}
