package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Get.json
func ExampleTransactionNodesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionNodesClient().Get(ctx, "contosemember1", "txnode2", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TransactionNode = armblockchain.TransactionNode{
	// 	Name: to.Ptr("txnode2"),
	// 	Type: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1/transactionNodes/txnode2"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armblockchain.TransactionNodeProperties{
	// 		DNS: to.Ptr("txnode2-contosemember1.blockchain.ppe.azure-int.net"),
	// 		ProvisioningState: to.Ptr(armblockchain.NodeProvisioningStateSucceeded),
	// 		PublicKey: to.Ptr("h7Q10I/1dLK/hzX8FkVrfl03D/aX8jW3YNoxJ/n4vkY="),
	// 		UserName: to.Ptr("txnode2"),
	// 	},
	// }
}
