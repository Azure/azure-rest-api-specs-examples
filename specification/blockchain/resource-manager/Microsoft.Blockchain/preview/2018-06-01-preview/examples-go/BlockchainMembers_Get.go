package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Get.json
func ExampleMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMembersClient().Get(ctx, "contosemember1", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Member = armblockchain.Member{
	// 	Name: to.Ptr("contosemember1"),
	// 	Type: to.Ptr("Microsoft.Blockchain/blockchainMembers"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armblockchain.MemberProperties{
	// 		Consortium: to.Ptr("ContoseConsortium"),
	// 		ConsortiumManagementAccountAddress: to.Ptr("0xc40d40dedc353885d3e8393735cf799e4d7fe38c"),
	// 		ConsortiumMemberDisplayName: to.Ptr("contosemember1"),
	// 		ConsortiumRole: to.Ptr("ADMIN"),
	// 		DNS: to.Ptr("contosemember1.blockchain.azure.com"),
	// 		ProvisioningState: to.Ptr(armblockchain.BlockchainMemberProvisioningStateSucceeded),
	// 		PublicKey: to.Ptr("1VhPX4PbNGnE9dOEjgTrw92dltBpKxFQjXWNugcwvl0="),
	// 		RootContractAddress: to.Ptr("0x7407947df2f67142340ca7d1a2c120f0dbfd30e1"),
	// 		UserName: to.Ptr("contosemember1"),
	// 		ValidatorNodesSKU: &armblockchain.MemberNodesSKU{
	// 			Capacity: to.Ptr[int32](2),
	// 		},
	// 		Protocol: to.Ptr(armblockchain.BlockchainProtocolQuorum),
	// 	},
	// }
}
