```go
package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Create.json
func ExampleMembersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contosemember1",
		"mygroup",
		&armblockchain.MembersClientBeginCreateOptions{BlockchainMember: &armblockchain.Member{
			Location: to.Ptr("southeastasia"),
			Properties: &armblockchain.MemberProperties{
				Consortium:                          to.Ptr("ContoseConsortium"),
				ConsortiumManagementAccountPassword: to.Ptr("<consortiumManagementAccountPassword>"),
				Password:                            to.Ptr("<password>"),
				ValidatorNodesSKU: &armblockchain.MemberNodesSKU{
					Capacity: to.Ptr[int32](2),
				},
				Protocol: to.Ptr(armblockchain.BlockchainProtocolQuorum),
			},
		},
		})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblockchain%2Farmblockchain%2Fv0.5.0/sdk/resourcemanager/blockchain/armblockchain/README.md) on how to add the SDK to your project and authenticate.
