package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetEntities.json
func ExampleEntitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEntitiesClient().NewListPager(&armmanagementgroups.EntitiesClientListOptions{Skiptoken: nil,
		Skip:         nil,
		Top:          nil,
		Select:       nil,
		Search:       nil,
		Filter:       nil,
		View:         nil,
		GroupName:    nil,
		CacheControl: nil,
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
		// page.EntityListResult = armmanagementgroups.EntityListResult{
		// 	Value: []*armmanagementgroups.EntityInfo{
		// 		{
		// 			Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Management/managementGroups"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
		// 			Properties: &armmanagementgroups.EntityInfoProperties{
		// 				DisplayName: to.Ptr("Group 1 Tenant 2"),
		// 				InheritedPermissions: to.Ptr(armmanagementgroups.PermissionsView),
		// 				NumberOfChildGroups: to.Ptr[int32](1),
		// 				NumberOfChildren: to.Ptr[int32](2),
		// 				NumberOfDescendants: to.Ptr[int32](3),
		// 				Parent: &armmanagementgroups.EntityParentGroupInfo{
		// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
		// 				},
		// 				ParentDisplayNameChain: []*string{
		// 					to.Ptr("Tenant Root Group"),
		// 					to.Ptr("Group 1 Tenant 2 Parent")},
		// 					ParentNameChain: []*string{
		// 						to.Ptr("TenantRootGroup"),
		// 						to.Ptr("Group1Tenant2Parent")},
		// 						Permissions: to.Ptr(armmanagementgroups.PermissionsView),
		// 						TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("20000000-0004-0000-0000-000000000000"),
		// 					Type: to.Ptr("Microsoft.Management/managementGroups"),
		// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0004-0000-0000-000000000000"),
		// 					Properties: &armmanagementgroups.EntityInfoProperties{
		// 						DisplayName: to.Ptr("Group 4 Tenant 2"),
		// 						InheritedPermissions: to.Ptr(armmanagementgroups.PermissionsDelete),
		// 						NumberOfChildGroups: to.Ptr[int32](0),
		// 						NumberOfChildren: to.Ptr[int32](0),
		// 						NumberOfDescendants: to.Ptr[int32](0),
		// 						Parent: &armmanagementgroups.EntityParentGroupInfo{
		// 							ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
		// 						},
		// 						ParentDisplayNameChain: []*string{
		// 							to.Ptr("Tenant Root Group"),
		// 							to.Ptr("Group 4 Tenant 2 Parent")},
		// 							ParentNameChain: []*string{
		// 								to.Ptr("TenantRootGroup"),
		// 								to.Ptr("Group4 Tenant2Parent")},
		// 								Permissions: to.Ptr(armmanagementgroups.PermissionsDelete),
		// 								TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 							},
		// 					}},
		// 				}
	}
}
