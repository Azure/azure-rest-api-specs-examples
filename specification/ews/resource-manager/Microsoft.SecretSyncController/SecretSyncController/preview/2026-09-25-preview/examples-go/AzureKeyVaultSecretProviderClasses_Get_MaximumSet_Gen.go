package armsecretsstoreextension_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/secretsstoreextension/armsecretsstoreextension"
)

// Generated from example definition: 2026-09-25-preview/AzureKeyVaultSecretProviderClasses_Get_MaximumSet_Gen.json
func ExampleAzureKeyVaultSecretProviderClassesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecretsstoreextension.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureKeyVaultSecretProviderClassesClient().Get(ctx, "rg-ssc-example", "akvspc-ssc-example", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecretsstoreextension.AzureKeyVaultSecretProviderClassesClientGetResponse{
	// 	AzureKeyVaultSecretProviderClass: armsecretsstoreextension.AzureKeyVaultSecretProviderClass{
	// 		Properties: &armsecretsstoreextension.AzureKeyVaultSecretProviderClassProperties{
	// 			ProvisioningState: to.Ptr(armsecretsstoreextension.ProvisioningStateSucceeded),
	// 			KeyvaultName: to.Ptr("example-ssc-key-vault"),
	// 			CloudName: to.Ptr(armsecretsstoreextension.AzureCloudNameAzurePublicCloud),
	// 			ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Objects: to.Ptr("array: |\n  - |\n    objectName: my-secret-object\n    objectType: secret\n    objectVersionHistory: 1"),
	// 		},
	// 		ExtendedLocation: &armsecretsstoreextension.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.ExtendedLocation/customLocations/example-custom-location"),
	// 			Type: to.Ptr(armsecretsstoreextension.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		Tags: map[string]*string{
	// 			"example-tag": to.Ptr("example-tag-value"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.SecretSyncController/azurekeyvaultsecretproviderclass/akvssc-example"),
	// 		Name: to.Ptr("akvspc-ssc-example"),
	// 	},
	// }
}
