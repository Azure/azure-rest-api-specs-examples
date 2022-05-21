Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv1.1.0-beta.1/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBMongoDBUserDefinitionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoUserDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewMongoDBResourcesClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateMongoUserDefinition(ctx,
		"myMongoUserDefinitionId",
		"myResourceGroupName",
		"myAccountName",
		armcosmos.MongoUserDefinitionCreateUpdateParameters{
			Properties: &armcosmos.MongoUserDefinitionResource{
				CustomData:   to.Ptr("My custom data"),
				DatabaseName: to.Ptr("sales"),
				Mechanisms:   to.Ptr("SCRAM-SHA-256"),
				Password:     to.Ptr("myPassword"),
				Roles: []*armcosmos.Role{
					{
						Db:   to.Ptr("sales"),
						Role: to.Ptr("myReadRole"),
					}},
				UserName: to.Ptr("myUserName"),
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
