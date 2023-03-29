package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListConsortiumMembers.json
func ExampleMembersClient_NewListConsortiumMembersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMembersClient().NewListConsortiumMembersPager("contosemember1", "mygroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ConsortiumMemberCollection = armblockchain.ConsortiumMemberCollection{
		// 	Value: []*armblockchain.ConsortiumMember{
		// 		{
		// 			Name: to.Ptr("contosemember1"),
		// 			DateModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-23T22:34:00.330Z"); return t}()),
		// 			DisplayName: to.Ptr("Contose member 1"),
		// 			JoinDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-23T22:34:00.330Z"); return t}()),
		// 			Role: to.Ptr("Member"),
		// 			Status: to.Ptr("Ready"),
		// 			SubscriptionID: to.Ptr("51766542-3ed7-4a72-a187-0c8ab644ddab"),
		// 		},
		// 		{
		// 			Name: to.Ptr(""),
		// 			DateModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-23T22:34:00.330Z"); return t}()),
		// 			Role: to.Ptr("ADMIN"),
		// 			Status: to.Ptr("Invited"),
		// 			SubscriptionID: to.Ptr("02bf808a-5446-4e54-aea8-39723aaa05e7"),
		// 	}},
		// }
	}
}
