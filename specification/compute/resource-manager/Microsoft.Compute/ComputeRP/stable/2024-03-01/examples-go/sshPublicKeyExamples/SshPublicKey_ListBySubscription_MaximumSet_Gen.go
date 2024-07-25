package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/sshPublicKeyExamples/SshPublicKey_ListBySubscription_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListBySubscriptionPager_sshPublicKeyListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListBySubscriptionPager(nil)
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
		// page.SSHPublicKeysGroupListResult = armcompute.SSHPublicKeysGroupListResult{
		// 	Value: []*armcompute.SSHPublicKeyResource{
		// 		{
		// 			Name: to.Ptr("mySshPublicKeyName"),
		// 			Type: to.Ptr("aaaa"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"key6396": to.Ptr("aaaaaaaaaaaaa"),
		// 				"key8839": to.Ptr("aaa"),
		// 			},
		// 			Properties: &armcompute.SSHPublicKeyResourceProperties{
		// 				PublicKey: to.Ptr("{ssh-rsa public key}"),
		// 			},
		// 	}},
		// }
	}
}
