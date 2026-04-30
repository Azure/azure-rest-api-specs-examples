package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListByResourceGroupPager_sshPublicKeyListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSSHPublicKeysClient().NewListByResourceGroupPager("rgcompute", nil)
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
		// page = armcompute.SSHPublicKeysClientListByResourceGroupResponse{
		// 	SSHPublicKeysGroupListResult: armcompute.SSHPublicKeysGroupListResult{
		// 		Value: []*armcompute.SSHPublicKeyResource{
		// 			{
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armcompute.SSHPublicKeyResourceProperties{
		// 					PublicKey: to.Ptr("{ssh-rsa public key}"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 				Name: to.Ptr("mySshPublicKeyName"),
		// 				Type: to.Ptr("aaaa"),
		// 				Tags: map[string]*string{
		// 					"key6396": to.Ptr("aaaaaaaaaaaaa"),
		// 					"key8839": to.Ptr("aaa"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaa"),
		// 	},
		// }
	}
}
