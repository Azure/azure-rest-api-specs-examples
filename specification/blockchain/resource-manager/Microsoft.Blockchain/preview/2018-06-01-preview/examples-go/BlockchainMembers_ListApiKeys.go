package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListApiKeys.json
func ExampleBlockchainMembersClient_ListAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	_, err = client.ListAPIKeys(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
