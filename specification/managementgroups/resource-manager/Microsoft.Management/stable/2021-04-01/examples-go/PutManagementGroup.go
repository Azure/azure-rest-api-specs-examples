package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"ChildGroup",
		armmanagementgroups.CreateManagementGroupRequest{
			Properties: &armmanagementgroups.CreateManagementGroupProperties{
				DisplayName: to.Ptr("ChildGroup"),
				Details: &armmanagementgroups.CreateManagementGroupDetails{
					Parent: &armmanagementgroups.CreateParentGroupInfo{
						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
					},
				},
			},
		},
		&armmanagementgroups.ClientBeginCreateOrUpdateOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
