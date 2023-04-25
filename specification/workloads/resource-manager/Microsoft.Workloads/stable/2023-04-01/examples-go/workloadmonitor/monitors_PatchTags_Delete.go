package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_PatchTags_Delete.json
func ExampleMonitorsClient_Update_deleteTagsFieldOfASapMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().Update(ctx, "myResourceGroup", "mySapMonitor", armworkloads.UpdateMonitorRequest{
		Identity: &armworkloads.UserAssignedServiceIdentity{
			Type: to.Ptr(armworkloads.ManagedServiceIdentityTypeNone),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Monitor = armworkloads.Monitor{
	// 	Name: to.Ptr("mySapMonitor"),
	// 	Type: to.Ptr("Microsoft.Workloads/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"testkey": to.Ptr("testvalue"),
	// 	},
	// 	Identity: &armworkloads.UserAssignedServiceIdentity{
	// 		Type: to.Ptr(armworkloads.ManagedServiceIdentityTypeNone),
	// 	},
	// 	Properties: &armworkloads.MonitorProperties{
	// 		AppLocation: to.Ptr("eastus"),
	// 		Errors: &armworkloads.MonitorPropertiesErrors{
	// 		},
	// 		LogAnalyticsWorkspaceArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.operationalinsights/workspaces/myWorkspace"),
	// 		ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
	// 			Name: to.Ptr("myManagedRg"),
	// 		},
	// 		MonitorSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 		MsiArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myMsi"),
	// 		ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
	// 		RoutingPreference: to.Ptr(armworkloads.RoutingPreferenceRouteAll),
	// 		StorageAccountArmID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/mystorageaccount"),
	// 		ZoneRedundancyPreference: to.Ptr("ZoneRedundantApp"),
	// 	},
	// }
}
