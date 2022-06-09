```go
package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBMongoDBRoleDefinitionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewMongoDBResourcesClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateMongoRoleDefinition(ctx,
		"myMongoRoleDefinitionId",
		"myResourceGroupName",
		"myAccountName",
		armcosmos.MongoRoleDefinitionCreateUpdateParameters{
			Properties: &armcosmos.MongoRoleDefinitionResource{
				DatabaseName: to.Ptr("sales"),
				Privileges: []*armcosmos.Privilege{
					{
						Actions: []*string{
							to.Ptr("insert"),
							to.Ptr("find")},
						Resource: &armcosmos.PrivilegeResource{
							Collection: to.Ptr("sales"),
							Db:         to.Ptr("sales"),
						},
					}},
				RoleName: to.Ptr("myRoleName"),
				Roles: []*armcosmos.Role{
					{
						Db:   to.Ptr("sales"),
						Role: to.Ptr("myInheritedRole"),
					}},
			},
		},
		nil)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv1.1.0-beta.1/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.
