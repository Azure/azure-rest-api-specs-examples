package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/sshPublicKeyExamples/SshPublicKey_Update_MaximumSet_Gen.json
func ExampleSSHPublicKeysClient_Update_sshPublicKeyUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSSHPublicKeysClient().Update(ctx, "rgcompute", "aaaaaaaaaaaa", armcompute.SSHPublicKeyUpdateResource{
		Tags: map[string]*string{
			"key2854": to.Ptr("a"),
		},
		Properties: &armcompute.SSHPublicKeyResourceProperties{
			PublicKey: to.Ptr("{ssh-rsa public key}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SSHPublicKeyResource = armcompute.SSHPublicKeyResource{
	// 	Name: to.Ptr("mySshPublicKeyName"),
	// 	Type: to.Ptr("aaaa"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/sshPublicKeys/mySshPublicKeyName"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key6396": to.Ptr("aaaaaaaaaaaaa"),
	// 		"key8839": to.Ptr("aaa"),
	// 	},
	// 	Properties: &armcompute.SSHPublicKeyResourceProperties{
	// 		PublicKey: to.Ptr("{ssh-rsa public key}"),
	// 	},
	// }
}
