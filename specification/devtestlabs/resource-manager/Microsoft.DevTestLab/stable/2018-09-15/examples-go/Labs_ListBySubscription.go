package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListBySubscription.json
func ExampleLabsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLabsClient().NewListBySubscriptionPager(&armdevtestlabs.LabsClientListBySubscriptionOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.LabList = armdevtestlabs.LabList{
		// 	Value: []*armdevtestlabs.Lab{
		// 		{
		// 			Name: to.Ptr("{labName1}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName1}"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.LabProperties{
		// 				Announcement: &armdevtestlabs.LabAnnouncementProperties{
		// 					Enabled: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					Expired: to.Ptr(false),
		// 					Markdown: to.Ptr(""),
		// 					Title: to.Ptr(""),
		// 				},
		// 				ArtifactsStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/{storageAccountName}"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-30T15:38:13.197Z"); return t}()),
		// 				DefaultPremiumStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/{storageAccountName}"),
		// 				DefaultStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/{storageAccountName}"),
		// 				EnvironmentPermission: to.Ptr(armdevtestlabs.EnvironmentPermissionReader),
		// 				LabStorageType: to.Ptr(armdevtestlabs.StorageTypePremium),
		// 				MandatoryArtifactsResourceIDsLinux: []*string{
		// 				},
		// 				MandatoryArtifactsResourceIDsWindows: []*string{
		// 				},
		// 				PremiumDataDiskStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/{storageAccountName}"),
		// 				PremiumDataDisks: to.Ptr(armdevtestlabs.PremiumDataDiskDisabled),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Support: &armdevtestlabs.LabSupportProperties{
		// 					Enabled: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					Markdown: to.Ptr(""),
		// 				},
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				VaultName: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.KeyVault/vaults/{keyVaultName}"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{labName2}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName2}/providers/microsoft.devtestlab/labs/{labName2}"),
		// 			Location: to.Ptr("japaneast"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.LabProperties{
		// 				Announcement: &armdevtestlabs.LabAnnouncementProperties{
		// 					Enabled: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					Expired: to.Ptr(false),
		// 					Markdown: to.Ptr(""),
		// 					Title: to.Ptr(""),
		// 				},
		// 				ArtifactsStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName2}/providers/Microsoft.Storage/storageAccounts/{storageAccountName2}"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-30T16:37:52.967Z"); return t}()),
		// 				DefaultPremiumStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName2}/providers/Microsoft.Storage/storageAccounts/{storageAccountName2}"),
		// 				DefaultStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName2}/providers/Microsoft.Storage/storageAccounts/{storageAccountName2}"),
		// 				EnvironmentPermission: to.Ptr(armdevtestlabs.EnvironmentPermissionReader),
		// 				LabStorageType: to.Ptr(armdevtestlabs.StorageTypePremium),
		// 				MandatoryArtifactsResourceIDsLinux: []*string{
		// 				},
		// 				MandatoryArtifactsResourceIDsWindows: []*string{
		// 				},
		// 				PremiumDataDiskStorageAccount: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName2}/providers/Microsoft.Storage/storageAccounts/{storageAccountName2}"),
		// 				PremiumDataDisks: to.Ptr(armdevtestlabs.PremiumDataDiskDisabled),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Support: &armdevtestlabs.LabSupportProperties{
		// 					Enabled: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					Markdown: to.Ptr(""),
		// 				},
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				VaultName: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName2}/providers/Microsoft.KeyVault/vaults/{keyVaultName2}"),
		// 			},
		// 	}},
		// }
	}
}
