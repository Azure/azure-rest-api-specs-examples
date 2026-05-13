package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups/v2"
)

// Generated from example definition: 2023-04-01/PatchManagementGroup.json
func ExampleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "ChildGroup", armmanagementgroups.PatchManagementGroupRequest{
		DisplayName:   to.Ptr("AlternateDisplayName"),
		ParentGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/AlternateRootGroup"),
	}, &armmanagementgroups.ClientUpdateOptions{
		CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmanagementgroups.ClientUpdateResponse{
	// 	ManagementGroup: armmanagementgroups.ManagementGroup{
	// 		Name: to.Ptr("ChildGroup"),
	// 		Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ChildGroup"),
	// 		Properties: &armmanagementgroups.ManagementGroupProperties{
	// 			DisplayName: to.Ptr("AlternateDisplayName"),
	// 			TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 			Details: &armmanagementgroups.ManagementGroupDetails{
	// 				Parent: &armmanagementgroups.ParentGroupInfo{
	// 					Name: to.Ptr("AlternateRootGroup"),
	// 					DisplayName: to.Ptr("AlternateRootGroup"),
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/AlternateRootGroup"),
	// 				},
	// 				UpdatedBy: to.Ptr("bd490e30-04cb-433e-b8c8-6066959a8bab"),
	// 				UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T02:46:59.0545645Z"); return t}()),
	// 				Version: to.Ptr[int32](2),
	// 			},
	// 		},
	// 	},
	// }
}
