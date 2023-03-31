package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceCreate.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginCreateOrUpdate_createOrUpdateAServiceWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().BeginCreateOrUpdate(ctx, "rg1", "service1", armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		Identity: &armm365securityandcompliance.ServicesResourceIdentity{
			Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
		Location: to.Ptr("westus2"),
		Tags:     map[string]*string{},
		Properties: &armm365securityandcompliance.ServicesProperties{
			AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
				{
					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				},
				{
					ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
				}},
			AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
				Audience:          to.Ptr("https://azurehealthcareapis.com"),
				Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
				SmartProxyEnabled: to.Ptr(true),
			},
			CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
				AllowCredentials: to.Ptr(false),
				Headers: []*string{
					to.Ptr("*")},
				MaxAge: to.Ptr[int64](1440),
				Methods: []*string{
					to.Ptr("DELETE"),
					to.Ptr("GET"),
					to.Ptr("OPTIONS"),
					to.Ptr("PATCH"),
					to.Ptr("POST"),
					to.Ptr("PUT")},
				Origins: []*string{
					to.Ptr("*")},
			},
			CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
				KeyVaultKeyURI:  to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
				OfferThroughput: to.Ptr[int64](1000),
			},
			ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
				StorageAccountName: to.Ptr("existingStorageAccount"),
			},
			PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
			PublicNetworkAccess:        to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
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
	// res.PrivateLinkServicesForM365SecurityCenterDescription = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
	// 	Identity: &armm365securityandcompliance.ServicesResourceIdentity{
	// 		Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("03fe6ae0-952c-4e4b-954b-cc0364dd252e"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d8cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armm365securityandcompliance.ServicesProperties{
	// 		AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 			},
	// 			{
	// 				ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
	// 		}},
	// 		AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 				to.Ptr("*")},
	// 				MaxAge: to.Ptr[int64](1440),
	// 				Methods: []*string{
	// 					to.Ptr("DELETE"),
	// 					to.Ptr("GET"),
	// 					to.Ptr("OPTIONS"),
	// 					to.Ptr("PATCH"),
	// 					to.Ptr("POST"),
	// 					to.Ptr("PUT")},
	// 					Origins: []*string{
	// 						to.Ptr("*")},
	// 					},
	// 					CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
	// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
	// 						OfferThroughput: to.Ptr[int64](1000),
	// 					},
	// 					ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
	// 						StorageAccountName: to.Ptr("existingStorageAccount"),
	// 					},
	// 					PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
	// 				},
	// 			}
}
