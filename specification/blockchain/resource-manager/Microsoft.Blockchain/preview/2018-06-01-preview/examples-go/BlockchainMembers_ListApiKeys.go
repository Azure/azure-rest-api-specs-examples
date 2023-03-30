package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListApiKeys.json
func ExampleMembersClient_ListAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMembersClient().ListAPIKeys(ctx, "contosemember1", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeyCollection = armblockchain.APIKeyCollection{
	// 	Keys: []*armblockchain.APIKey{
	// 		{
	// 			KeyName: to.Ptr("key1"),
	// 			Value: to.Ptr("-EnIUzu29xj60xPJmstyCURo"),
	// 		},
	// 		{
	// 			KeyName: to.Ptr("key2"),
	// 			Value: to.Ptr("I8P-q4u_WDuCZUBYXnJ3yYX7"),
	// 	}},
	// }
}
