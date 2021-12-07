Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblockchain%2Farmblockchain%2Fv0.1.0/sdk/resourcemanager/blockchain/armblockchain/README.md) on how to add the SDK to your project and authenticate.

```go
package armblockchain_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Create.json
func ExampleBlockchainMembersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		&armblockchain.BlockchainMembersBeginCreateOptions{BlockchainMember: &armblockchain.BlockchainMember{
			TrackedResource: armblockchain.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armblockchain.BlockchainMemberProperties{
				Consortium:                          to.StringPtr("<consortium>"),
				ConsortiumManagementAccountPassword: to.StringPtr("<consortium-management-account-password>"),
				Password:                            to.StringPtr("<password>"),
				ValidatorNodesSKU: &armblockchain.BlockchainMemberNodesSKU{
					Capacity: to.Int32Ptr(2),
				},
				Protocol: armblockchain.BlockchainProtocolQuorum.ToPtr(),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockchainMember.ID: %s\n", *res.ID)
}
```
