package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresList.json
func ExampleConfigurationStoresClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationStoresClient().NewListPager(&armappconfiguration.ConfigurationStoresClientListOptions{SkipToken: nil})
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
		// page.ConfigurationStoreListResult = armappconfiguration.ConfigurationStoreListResult{
		// 	Value: []*armappconfiguration.ConfigurationStore{
		// 		{
		// 			Name: to.Ptr("contoso"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armappconfiguration.ResourceIdentity{
		// 				Type: to.Ptr(armappconfiguration.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA"),
		// 				TenantID: to.Ptr("BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"),
		// 			},
		// 			Properties: &armappconfiguration.ConfigurationStoreProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				Encryption: &armappconfiguration.EncryptionProperties{
		// 					KeyVaultProperties: &armappconfiguration.KeyVaultProperties{
		// 					},
		// 				},
		// 				Endpoint: to.Ptr("https://contoso.azconfig.io"),
		// 				PrivateEndpointConnections: []*armappconfiguration.PrivateEndpointConnectionReference{
		// 				},
		// 				ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armappconfiguration.PublicNetworkAccessDisabled),
		// 			},
		// 			SKU: &armappconfiguration.SKU{
		// 				Name: to.Ptr("Standard"),
		// 			},
		// 			SystemData: &armappconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 				CreatedBy: to.Ptr("foo@contoso.com"),
		// 				CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("contoso2"),
		// 			Type: to.Ptr("Microsoft.AppConfiguration/configurationStores"),
		// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.AppConfiguration/configurationStores/contoso2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armappconfiguration.ResourceIdentity{
		// 				Type: to.Ptr(armappconfiguration.IdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC"),
		// 				TenantID: to.Ptr("BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"),
		// 			},
		// 			Properties: &armappconfiguration.ConfigurationStoreProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T23:06:59+00:00"); return t}()),
		// 				DisableLocalAuth: to.Ptr(false),
		// 				Encryption: &armappconfiguration.EncryptionProperties{
		// 					KeyVaultProperties: &armappconfiguration.KeyVaultProperties{
		// 					},
		// 				},
		// 				Endpoint: to.Ptr("https://contoso2.azconfig.io"),
		// 				PrivateEndpointConnections: []*armappconfiguration.PrivateEndpointConnectionReference{
		// 				},
		// 				ProvisioningState: to.Ptr(armappconfiguration.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armappconfiguration.PublicNetworkAccessDisabled),
		// 			},
		// 			SKU: &armappconfiguration.SKU{
		// 				Name: to.Ptr("Standard"),
		// 			},
		// 			SystemData: &armappconfiguration.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 				CreatedBy: to.Ptr("foo@contoso.com"),
		// 				CreatedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
		// 				LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armappconfiguration.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
