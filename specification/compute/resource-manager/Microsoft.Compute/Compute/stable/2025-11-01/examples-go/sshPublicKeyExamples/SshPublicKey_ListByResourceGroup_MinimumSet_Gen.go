package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MinimumSet_Gen.json
func ExampleSSHPublicKeysClient_NewListByResourceGroupPager_sshPublicKeyListByResourceGroupMinimumSetGen() {
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
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
