package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview/v2"
)

// Generated from example definition: 2024-04-01-preview/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("12345678-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "SampleResourceGroup", "account1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurview.AccountsClientGetResponse{
	// 	Account: &armpurview.Account{
	// 		Name: to.Ptr("account1"),
	// 		Type: to.Ptr("Microsoft.Purview/accounts"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
	// 		Location: to.Ptr("West US 2"),
	// 		Properties: &armpurview.AccountProperties{
	// 			AccountStatus: &armpurview.AccountPropertiesAccountStatus{
	// 				AccountProvisioningState: to.Ptr(armpurview.AccountProvisioningStateSucceeded),
	// 				ErrorDetails: &armpurview.AccountStatusErrorDetails{
	// 				},
	// 			},
	// 			DefaultDomain: to.Ptr("account1"),
	// 			Endpoints: &armpurview.AccountPropertiesEndpoints{
	// 				Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
	// 				Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
	// 			},
	// 			FriendlyName: to.Ptr("friendly-account1"),
	// 			IngestionStorage: &armpurview.IngestionStorage{
	// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/ingestion-storage-rg-mwjotkl"),
	// 				PrimaryEndpoint: to.Ptr("https://ingestioneastusddnpkar.z31.blob.storage.azure.net/"),
	// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 			},
	// 			ManagedResourceGroupName: to.Ptr("managed-rg-mwjotkl"),
	// 			ManagedResources: &armpurview.AccountPropertiesManagedResources{
	// 				EventHubNamespace: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.EventHub/namespaces/atlas-westusdddnbtp"),
	// 				ResourceGroup: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl"),
	// 				StorageAccount: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.Storage/storageAccounts/scanwestustzaagzr"),
	// 			},
	// 			ManagedResourcesPublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 			MergeInfo: &armpurview.AccountMergeInfo{
	// 				AccountLocation: to.Ptr("westus2"),
	// 				AccountName: to.Ptr("otherAccount"),
	// 				AccountResourceGroupName: to.Ptr("otherAccountRG"),
	// 				AccountSubscriptionID: to.Ptr("92d4ce06-0032-4e6b-9103-88ba3e159551"),
	// 				Deprovisioned: to.Ptr(false),
	// 				MergeStatus: to.Ptr(armpurview.MergeStatusSucceeded),
	// 				TypeOfAccount: to.Ptr(armpurview.MergeAccountTypeSecondary),
	// 			},
	// 			PrivateEndpointConnections: []*armpurview.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 					Type: to.Ptr("Microsoft.Purview/accounts/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateEndpointConnections/peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 					Properties: &armpurview.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armpurview.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/baca8a88-4527-4c35-a13e-b2775ce0d7fc/resourceGroups/nrpResourceGroupName/providers/Microsoft.Network/privateEndpoints/peName"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Please approve my connection, thanks."),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armpurview.PrivateEndpointConnectionStatusPending),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 			TenantEndpointState: to.Ptr(armpurview.TenantEndpointStateDisabled),
	// 		},
	// 		SKU: &armpurview.AccountSKU{
	// 			Name: to.Ptr(armpurview.AccountSKUNameStandard),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armpurview.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 			CreatedBy: to.Ptr("client-name"),
	// 			CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("client-name"),
	// 			LastModifiedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
