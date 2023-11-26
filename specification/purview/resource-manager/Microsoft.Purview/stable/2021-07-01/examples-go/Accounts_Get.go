package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.Account = armpurview.Account{
	// 	Name: to.Ptr("account1"),
	// 	Type: to.Ptr("Microsoft.Purview/accounts"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armpurview.TrackedResourceSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.343Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("client-name"),
	// 		LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armpurview.AccountProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByObjectID: to.Ptr("client-objectId"),
	// 		Endpoints: &armpurview.AccountPropertiesEndpoints{
	// 			Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
	// 			Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
	// 			Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
	// 		},
	// 		FriendlyName: to.Ptr("friendly-account1"),
	// 		ManagedResourceGroupName: to.Ptr("managed-rg-mwjotkl"),
	// 		ManagedResources: &armpurview.AccountPropertiesManagedResources{
	// 			EventHubNamespace: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.EventHub/namespaces/atlas-westusdddnbtp"),
	// 			ResourceGroup: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl"),
	// 			StorageAccount: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.Storage/storageAccounts/scanwestustzaagzr"),
	// 		},
	// 		PrivateEndpointConnections: []*armpurview.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 				Type: to.Ptr("Microsoft.Purview/accounts/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateEndpointConnections/peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 				Properties: &armpurview.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armpurview.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/baca8a88-4527-4c35-a13e-b2775ce0d7fc/resourceGroups/nrpResourceGroupName/providers/Microsoft.Network/privateEndpoints/peName"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my connection, thanks."),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armpurview.StatusPending),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armpurview.AccountSKU{
	// 		Name: to.Ptr(armpurview.NameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// }
}
