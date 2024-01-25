package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDataProductsClient_NewListByResourceGroupPager_dataProductsListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataProductsClient().NewListByResourceGroupPager("aoiresourceGroupName", nil)
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
		// page.DataProductListResult = armnetworkanalytics.DataProductListResult{
		// 	Value: []*armnetworkanalytics.DataProduct{
		// 		{
		// 			Name: to.Ptr("dataproduct01"),
		// 			Type: to.Ptr("Microsoft.NetworkAnalytics/DataProducts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/aoiresourceGroupName/providers/Microsoft.NetworkAnalytics/DataProducts/dataproduct01"),
		// 			SystemData: &armnetworkanalytics.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
		// 				CreatedBy: to.Ptr("abc@micros.com"),
		// 				CreatedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("abc@micros.com"),
		// 				LastModifiedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"userSpecifiedKeyName": to.Ptr("userSpecifiedKeyValue"),
		// 			},
		// 			Identity: &armnetworkanalytics.ManagedServiceIdentity{
		// 				Type: to.Ptr(armnetworkanalytics.ManagedServiceIdentityType("IdentityType")),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				UserAssignedIdentities: map[string]*armnetworkanalytics.UserAssignedIdentity{
		// 					"key8474": &armnetworkanalytics.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armnetworkanalytics.DataProductProperties{
		// 				AvailableMinorVersions: []*string{
		// 					to.Ptr("1.0.1"),
		// 					to.Ptr("1.0.2")},
		// 					ConsumptionEndpoints: &armnetworkanalytics.ConsumptionEndpointsProperties{
		// 						FileAccessResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/storageResourceName"),
		// 						FileAccessURL: to.Ptr("https://operatorinsightsstorageResourceName.blob.core.windows.net"),
		// 						IngestionResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.Storage/storageAccounts/storageResourceName"),
		// 						IngestionURL: to.Ptr("https://aoiingestionstorageResourceName.blob.core.windows.net"),
		// 						QueryResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.Kusto/clusters/clusterName"),
		// 						QueryURL: to.Ptr("https://opinsightsclusterName.regionName.kusto.windows.net"),
		// 					},
		// 					CurrentMinorVersion: to.Ptr("1.0.1"),
		// 					CustomerEncryptionKey: &armnetworkanalytics.EncryptionKeyDetails{
		// 						KeyName: to.Ptr("keyName"),
		// 						KeyVaultURI: to.Ptr("https://KeyVault.vault.azure.net"),
		// 						KeyVersion: to.Ptr("keyVersion"),
		// 					},
		// 					CustomerManagedKeyEncryptionEnabled: to.Ptr(armnetworkanalytics.ControlStateEnabled),
		// 					Documentation: to.Ptr("https://learn.microsoft.com/"),
		// 					KeyVaultURL: to.Ptr("https://myKeyVault.vault.azure.net"),
		// 					MajorVersion: to.Ptr("1.0.0"),
		// 					ManagedResourceGroupConfiguration: &armnetworkanalytics.ManagedResourceGroupConfiguration{
		// 						Name: to.Ptr("managedResourceGroupName"),
		// 						Location: to.Ptr("eastus"),
		// 					},
		// 					Networkacls: &armnetworkanalytics.DataProductNetworkACLs{
		// 						AllowedQueryIPRangeList: []*string{
		// 							to.Ptr("1.1.1.1"),
		// 							to.Ptr("1.1.1.2")},
		// 							DefaultAction: to.Ptr(armnetworkanalytics.DefaultActionAllow),
		// 							IPRules: []*armnetworkanalytics.IPRules{
		// 								{
		// 									Action: to.Ptr("Allow"),
		// 									Value: to.Ptr("1.1.1.1"),
		// 							}},
		// 							VirtualNetworkRule: []*armnetworkanalytics.VirtualNetworkRule{
		// 								{
		// 									Action: to.Ptr("Allow"),
		// 									ID: to.Ptr("/subscriptions/subscriptionId/resourcegroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/virtualNetworkName/subnets/subnetName"),
		// 									State: to.Ptr("Succeeded"),
		// 							}},
		// 						},
		// 						Owners: []*string{
		// 							to.Ptr("abc@micros.com")},
		// 							PrivateLinksEnabled: to.Ptr(armnetworkanalytics.ControlStateDisabled),
		// 							Product: to.Ptr("MCC"),
		// 							ProvisioningState: to.Ptr(armnetworkanalytics.ProvisioningStateSucceeded),
		// 							PublicNetworkAccess: to.Ptr(armnetworkanalytics.ControlStateEnabled),
		// 							Publisher: to.Ptr("Microsoft"),
		// 							PurviewAccount: to.Ptr("testpurview"),
		// 							PurviewCollection: to.Ptr("134567890"),
		// 							Redundancy: to.Ptr(armnetworkanalytics.ControlStateDisabled),
		// 							ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						},
		// 				}},
		// 			}
	}
}
