package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: 2025-06-01-preview/ConfigurationStoresCreateWithAzureFrontDoor.json
func ExampleConfigurationStoresClient_BeginCreate_configurationStoresCreateWithAzureFrontDoor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationStoresClient().BeginCreate(ctx, "myResourceGroup", "contoso", armappconfiguration.ConfigurationStore{
		Location: to.Ptr("westus"),
		SKU: &armappconfiguration.SKU{
			Name: to.Ptr("Standard"),
		},
		Tags: map[string]*string{
			"myTag": to.Ptr("myTagValue"),
		},
		Properties: &armappconfiguration.ConfigurationStoreProperties{
			AzureFrontDoor: &armappconfiguration.AzureFrontDoorProperties{
				ResourceID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/microsoft.cdn/profiles/myAzureFrontDoorProfile"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappconfiguration.ConfigurationStoresClientCreateResponse{
	// 	ConfigurationStore: &armappconfiguration.ConfigurationStore{
	// 		Type: to.Ptr("Microsoft.AppConfiguration/configurationStores"),
	// 		Properties: &armappconfiguration.ConfigurationStoreProperties{
	// 			ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			Endpoint: to.Ptr("https://contoso.azconfig.io"),
	// 			Encryption: &armappconfiguration.EncryptionProperties{
	// 				KeyVaultProperties: &armappconfiguration.KeyVaultProperties{
	// 				},
	// 			},
	// 			DisableLocalAuth: to.Ptr(false),
	// 			PrivateEndpointConnections: []*armappconfiguration.PrivateEndpointConnectionReference{
	// 			},
	// 			SoftDeleteRetentionInDays: to.Ptr[int32](30),
	// 			EnablePurgeProtection: to.Ptr(false),
	// 			DataPlaneProxy: &armappconfiguration.DataPlaneProxyProperties{
	// 				AuthenticationMode: to.Ptr(armappconfiguration.AuthenticationModeLocal),
	// 				PrivateLinkDelegation: to.Ptr(armappconfiguration.PrivateLinkDelegationDisabled),
	// 			},
	// 			AzureFrontDoor: &armappconfiguration.AzureFrontDoorProperties{
	// 				ResourceID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/microsoft.cdn/profiles/myAzureFrontDoorProfile"),
	// 			},
	// 		},
	// 		SystemData: &armappconfiguration.SystemData{
	// 			CreatedBy: to.Ptr("foo@contoso.com"),
	// 			CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 		},
	// 		SKU: &armappconfiguration.SKU{
	// 			Name: to.Ptr("Standard"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso"),
	// 		Name: to.Ptr("contoso"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag": to.Ptr("myTagValue"),
	// 		},
	// 	},
	// }
}
