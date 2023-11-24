package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithPath.json
func ExampleClient_Get_getManagementGroupWithPath() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", &armmanagementgroups.ClientGetOptions{Expand: to.Ptr(armmanagementgroups.ManagementGroupExpandTypePath),
		Recurse:      nil,
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("Group 1 Tenant 2"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Path: []*armmanagementgroups.ManagementGroupPathElement{
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 					DisplayName: to.Ptr("Grandparent display name"),
	// 				},
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 					DisplayName: to.Ptr("Parent display name"),
	// 			}},
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 				DisplayName: to.Ptr("Parent display name"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000001"),
	// 			},
	// 			UpdatedBy: to.Ptr("Test"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}
