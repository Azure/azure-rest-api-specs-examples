package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/ConfigurationStoresGet.json
func ExampleConfigurationStoresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationStoresClient().Get(ctx, "myResourceGroup", "contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 	Identity: &armappconfiguration.ResourceIdentity{
	// 		Type: to.Ptr(armappconfiguration.IdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA"),
	// 		TenantID: to.Ptr("BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"),
	// 	},
	// 	Properties: &armappconfiguration.ConfigurationStoreProperties{
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
	// 		DataPlaneProxy: &armappconfiguration.DataPlaneProxyProperties{
	// 			AuthenticationMode: to.Ptr(armappconfiguration.AuthenticationModePassThrough),
	// 			PrivateLinkDelegation: to.Ptr(armappconfiguration.PrivateLinkDelegationDisabled),
	// 		},
	// 		DefaultKeyValueRevisionRetentionPeriodInSeconds: to.Ptr[int64](2592000),
	// 		DisableLocalAuth: to.Ptr(false),
	// 		EnablePurgeProtection: to.Ptr(false),
	// 		Encryption: &armappconfiguration.EncryptionProperties{
	// 			KeyVaultProperties: &armappconfiguration.KeyVaultProperties{
	// 			},
	// 		},
	// 		Endpoint: to.Ptr("https://contoso.azconfig.io"),
	// 		PrivateEndpointConnections: []*armappconfiguration.PrivateEndpointConnectionReference{
	// 		},
	// 		ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armappconfiguration.PublicNetworkAccessDisabled),
	// 		SoftDeleteRetentionInDays: to.Ptr[int32](30),
	// 	},
	// 	SKU: &armappconfiguration.SKU{
	// 		Name: to.Ptr("Standard"),
	// 	},
	// 	SystemData: &armappconfiguration.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("foo@contoso.com"),
	// 		CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
	// 	},
	// }
}
