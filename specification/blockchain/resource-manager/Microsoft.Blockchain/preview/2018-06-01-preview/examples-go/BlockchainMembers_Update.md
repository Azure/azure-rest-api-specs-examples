Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblockchain%2Farmblockchain%2Fv0.1.0/sdk/resourcemanager/blockchain/armblockchain/README.md) on how to add the SDK to your project and authenticate.

```go
package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Update.json
func ExampleBlockchainMembersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		&armblockchain.BlockchainMembersUpdateOptions{BlockchainMember: &armblockchain.BlockchainMemberUpdate{
			Properties: &armblockchain.BlockchainMemberPropertiesUpdate{
				TransactionNodePropertiesUpdate: armblockchain.TransactionNodePropertiesUpdate{
					FirewallRules: []*armblockchain.FirewallRule{},
					Password:      to.StringPtr("<password>"),
				},
				ConsortiumManagementAccountPassword: to.StringPtr("<consortium-management-account-password>"),
			},
			Tags: map[string]*string{},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockchainMember.ID: %s\n", *res.ID)
}
```
