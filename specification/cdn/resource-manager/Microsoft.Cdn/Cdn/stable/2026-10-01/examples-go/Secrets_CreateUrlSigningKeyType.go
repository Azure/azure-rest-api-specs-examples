package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v4"
)

// Generated from example definition: 2026-10-01/Secrets_CreateUrlSigningKeyType.json
func ExampleSecretsClient_BeginCreate_secretsCreateUrlSigningKeyType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecretsClient().BeginCreate(ctx, "RG", "profile1", "secret1", armcdn.Secret{
		Properties: &armcdn.SecretProperties{
			Parameters: &armcdn.URLSigningKeyParameters{
				Type:  to.Ptr(armcdn.SecretTypeURLSigningKey),
				KeyID: to.Ptr("customKeyId"),
				SecretSource: &armcdn.ResourceReference{
					ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/urlsigningkeyname"),
				},
				SecretVersion: to.Ptr("abcdef1234578900abcdef1234567890"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcdn.SecretsClientCreateResponse{
	// 	Secret: armcdn.Secret{
	// 		Name: to.Ptr("secret1"),
	// 		Type: to.Ptr("Microsoft.Cdn/profiles/secrets"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/secrets/secret1"),
	// 		Properties: &armcdn.SecretProperties{
	// 			DeploymentStatus: to.Ptr(armcdn.DeploymentStatusNotStarted),
	// 			Parameters: &armcdn.URLSigningKeyParameters{
	// 				Type: to.Ptr(armcdn.SecretTypeURLSigningKey),
	// 				KeyID: to.Ptr("customKeyId"),
	// 				SecretSource: &armcdn.ResourceReference{
	// 					ID: to.Ptr("/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/urlsigningkeyname"),
	// 				},
	// 				SecretVersion: to.Ptr("abcdef1234578900abcdef1234567890"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcdn.AfdProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
