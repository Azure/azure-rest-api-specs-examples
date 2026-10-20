package armsecretsstoreextension_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/secretsstoreextension/armsecretsstoreextension"
)

// Generated from example definition: 2026-09-25-preview/AzureKeyVaultSecretProviderClasses_ListBySubscription_MaximumSet_Gen.json
func ExampleAzureKeyVaultSecretProviderClassesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecretsstoreextension.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureKeyVaultSecretProviderClassesClient().NewListBySubscriptionPager(nil)
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
		// page = armsecretsstoreextension.AzureKeyVaultSecretProviderClassesClientListBySubscriptionResponse{
		// 	AzureKeyVaultSecretProviderClassListResult: armsecretsstoreextension.AzureKeyVaultSecretProviderClassListResult{
		// 		Value: []*armsecretsstoreextension.AzureKeyVaultSecretProviderClass{
		// 			{
		// 				Properties: &armsecretsstoreextension.AzureKeyVaultSecretProviderClassProperties{
		// 					ProvisioningState: to.Ptr(armsecretsstoreextension.ProvisioningStateSucceeded),
		// 					KeyvaultName: to.Ptr("example-ssc-key-vault"),
		// 					CloudName: to.Ptr(armsecretsstoreextension.AzureCloudNameAzurePublicCloud),
		// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Objects: to.Ptr("array: |\n  - |\n    objectName: my-secret-object\n    objectType: secret\n    objectVersionHistory: 1"),
		// 				},
		// 				ExtendedLocation: &armsecretsstoreextension.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.ExtendedLocation/customLocations/example-custom-location"),
		// 					Type: to.Ptr(armsecretsstoreextension.ExtendedLocationTypeCustomLocation),
		// 				},
		// 				Tags: map[string]*string{
		// 					"example-tag": to.Ptr("example-tag-value"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.SecretSyncController/azurekeyvaultsecretproviderclass/akvssc-example"),
		// 				Name: to.Ptr("akvspc-ssc-example"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aolkmb"),
		// 	},
		// }
	}
}
