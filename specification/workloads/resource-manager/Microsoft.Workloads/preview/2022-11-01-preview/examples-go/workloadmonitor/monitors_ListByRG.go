package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/monitors_ListByRG.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewMonitorsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("example-rg", nil)
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
		// 			Type: to.Ptr("Microsoft.Workloads/workloads"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Workloads/monitors/mySapMonitor1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armworkloads.MonitorProperties{
		// 				AppLocation: to.Ptr("eastus"),
		// 				Errors: &armworkloads.MonitorPropertiesErrors{
		// 				},
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/example-rg/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
		// 					Name: to.Ptr("myManagedRg1"),
		// 				},
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				MsiArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myMsi"),
		// 				ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				RoutingPreference: to.Ptr(armworkloads.RoutingPreferenceRouteAll),
		// 				StorageAccountArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		// 				ZoneRedundancyPreference: to.Ptr("ZoneRedundantApp"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySapMonitor2"),
		// 			Type: to.Ptr("Microsoft.Workloads/workloads"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Workloads/workloads/mySapMonitor2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armworkloads.MonitorProperties{
		// 				AppLocation: to.Ptr("westus"),
		// 				Errors: &armworkloads.MonitorPropertiesErrors{
		// 				},
		// 				LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/example-rg/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
		// 				ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
		// 					Name: to.Ptr("myManagedRg2"),
		// 				},
		// 				MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				MsiArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myMsi"),
		// 				ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				RoutingPreference: to.Ptr(armworkloads.RoutingPreferenceDefault),
		// 				StorageAccountArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
		// 			},
		// 	}},
		// }
	}
}
