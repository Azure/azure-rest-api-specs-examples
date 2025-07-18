package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/UpdateVaultExtendedInfo.json
func ExampleVaultExtendedInfoClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVaultExtendedInfoClient().Update(ctx, "Default-RecoveryServices-ResourceGroup", "swaggerExample", armrecoveryservices.VaultExtendedInfoResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VaultExtendedInfoResource = armrecoveryservices.VaultExtendedInfoResource{
	// 	Name: to.Ptr("vaultExtendedInfo"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/extendedInformation"),
	// 	Etag: to.Ptr("f0d0260b-b92d-4458-ba0a-32c6cdabacb7"),
	// 	ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/resourceGroups/Default-RecoveryServices-ResourceGroup/providers/Microsoft.RecoveryServices/vaults/swaggerExample/extendedInformation/vaultExtendedInfo"),
	// 	Properties: &armrecoveryservices.VaultExtendedInfo{
	// 		Algorithm: to.Ptr("None"),
	// 		IntegrityKey: to.Ptr("J99wzS27fmJ+Wjot7xO5wA=="),
	// 	},
	// }
}
