package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutHierarchySettings.json
func ExampleHierarchySettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHierarchySettingsClient().CreateOrUpdate(ctx, "root", armmanagementgroups.CreateOrUpdateSettingsRequest{
		Properties: &armmanagementgroups.CreateOrUpdateSettingsProperties{
			DefaultManagementGroup:               to.Ptr("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
			RequireAuthorizationForGroupCreation: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HierarchySettings = armmanagementgroups.HierarchySettings{
	// 	Name: to.Ptr("root"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups/settings"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/root/settings/default"),
	// 	Properties: &armmanagementgroups.HierarchySettingsProperties{
	// 		DefaultManagementGroup: to.Ptr("/providers/Microsoft.Management/managementGroups/DefaultGroup"),
	// 		RequireAuthorizationForGroupCreation: to.Ptr(true),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 	},
	// }
}
