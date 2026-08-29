package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-05-01-preview/StoragePools_ConfigurePlatformConsoleAuth_MaximumSet_Gen.json
func ExampleStoragePoolsClient_ConfigurePlatformConsoleAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStoragePoolsClient().ConfigurePlatformConsoleAuth(ctx, "rgpurestorage", "storagepool-01", &armpurestorageblock.SSHPlatformConsoleAuthConfig{
		AuthType:  to.Ptr(armpurestorageblock.PlatformConsoleAuthTypeSSH),
		Username:  to.Ptr("alice"),
		PublicKey: to.Ptr("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl alice@example.com"),
		Role:      to.Ptr(armpurestorageblock.PlatformConsoleRoleStorageAdmin),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.StoragePoolsClientConfigurePlatformConsoleAuthResponse{
	// 	PlatformConsoleAuthResultClassification: &armpurestorageblock.SSHPlatformConsoleAuthResult{
	// 		AuthType: to.Ptr(armpurestorageblock.PlatformConsoleAuthTypeSSH),
	// 		Username: to.Ptr("alice"),
	// 		Role: to.Ptr(armpurestorageblock.PlatformConsoleRoleStorageAdmin),
	// 	},
	// }
}
