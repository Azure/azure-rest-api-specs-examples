package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "ChildGroup", armmanagementgroups.CreateManagementGroupRequest{
		Properties: &armmanagementgroups.CreateManagementGroupProperties{
			DisplayName: to.Ptr("ChildGroup"),
			Details: &armmanagementgroups.CreateManagementGroupDetails{
				Parent: &armmanagementgroups.CreateParentGroupInfo{
					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
				},
			},
		},
	}, &armmanagementgroups.ClientBeginCreateOrUpdateOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("ChildGroup"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ChildGroup"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("ChildGroup"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("RootGroup"),
	// 				DisplayName: to.Ptr("RootGroup"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
	// 			},
	// 			UpdatedBy: to.Ptr("16b8ef21-5c9f-420c-bcc9-e4f8c9f30b4b"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}
