package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListByResourceGroup.json
func ExampleLabsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLabsClient().NewListByResourceGroupPager("resourceGroupName", &armdevtestlabs.LabsClientListByResourceGroupOptions{Expand: nil,
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
		// 			Name: to.Ptr("{labName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
		// 			Location: to.Ptr("{location}"),
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
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-02T17:23:20.957Z"); return t}()),
		// 				EnvironmentPermission: to.Ptr(armdevtestlabs.EnvironmentPermissionReader),
		// 				LabStorageType: to.Ptr(armdevtestlabs.StorageTypeStandard),
		// 				MandatoryArtifactsResourceIDsLinux: []*string{
		// 				},
		// 				MandatoryArtifactsResourceIDsWindows: []*string{
		// 				},
		// 				PremiumDataDisks: to.Ptr(armdevtestlabs.PremiumDataDiskDisabled),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Support: &armdevtestlabs.LabSupportProperties{
		// 					Enabled: to.Ptr(armdevtestlabs.EnableStatusDisabled),
		// 					Markdown: to.Ptr(""),
		// 				},
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				VaultName: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.KeyVault/vaults/{keyVaultName}"),
		// 			},
		// 	}},
		// }
	}
}
