package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresUpdateDisableLocalAuth.json
func ExampleConfigurationStoresClient_BeginUpdate_configurationStoresUpdateDisableLocalAuth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationStoresClient().BeginUpdate(ctx, "myResourceGroup", "contoso", armappconfiguration.ConfigurationStoreUpdateParameters{
		Properties: &armappconfiguration.ConfigurationStorePropertiesUpdateParameters{
			DisableLocalAuth: to.Ptr(true),
		},
		SKU: &armappconfiguration.SKU{
			Name: to.Ptr("Standard"),
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
	// res.ConfigurationStore = armappconfiguration.ConfigurationStore{
	// 	Name: to.Ptr("contoso"),
	// 	Type: to.Ptr("Microsoft.AppConfiguration/configurationStores"),
	// 	ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armappconfiguration.ConfigurationStoreProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 		DisableLocalAuth: to.Ptr(true),
	// 		Encryption: &armappconfiguration.EncryptionProperties{
	// 			KeyVaultProperties: &armappconfiguration.KeyVaultProperties{
	// 			},
	// 		},
	// 		Endpoint: to.Ptr("https://contoso.azconfig.io"),
	// 		PrivateEndpointConnections: []*armappconfiguration.PrivateEndpointConnectionReference{
	// 		},
	// 		ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armappconfiguration.PublicNetworkAccessDisabled),
	// 	},
	// 	SKU: &armappconfiguration.SKU{
	// 		Name: to.Ptr("Standard"),
	// 	},
	// 	SystemData: &armappconfiguration.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 		CreatedBy: to.Ptr("foo@contoso.com"),
	// 		CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 		LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 	},
	// }
}
