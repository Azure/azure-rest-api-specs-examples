package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/ServiceCreate.json
func ExampleServicesClient_BeginCreateOrUpdate_createOrUpdateAServiceWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "rg1", "service1", armhealthcareapis.ServicesDescription{
		Identity: &armhealthcareapis.ServicesResourceIdentity{
			Type: to.Ptr(armhealthcareapis.ManagedServiceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr(armhealthcareapis.KindFhirR4),
		Location: to.Ptr("westus2"),
		Tags:     map[string]*string{},
		Properties: &armhealthcareapis.ServicesProperties{
			AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
				{
					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				},
				{
					ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
				}},
			AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
				Audience:          to.Ptr("https://azurehealthcareapis.com"),
				Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
				SmartProxyEnabled: to.Ptr(true),
			},
			CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
				AllowCredentials: to.Ptr(false),
				Headers: []*string{
					to.Ptr("*")},
				MaxAge: to.Ptr[int32](1440),
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
			CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
				KeyVaultKeyURI:  to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
				OfferThroughput: to.Ptr[int32](1000),
			},
			ExportConfiguration: &armhealthcareapis.ServiceExportConfigurationInfo{
				StorageAccountName: to.Ptr("existingStorageAccount"),
			},
			PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{},
			PublicNetworkAccess:        to.Ptr(armhealthcareapis.PublicNetworkAccessDisabled),
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
	// res.ServicesDescription = armhealthcareapis.ServicesDescription{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/services"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/services/service1"),
	// 	Identity: &armhealthcareapis.ServicesResourceIdentity{
	// 		Type: to.Ptr(armhealthcareapis.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("03fe6ae0-952c-4e4b-954b-cc0364dd252e"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d8cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armhealthcareapis.KindFhirR4),
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armhealthcareapis.ServicesProperties{
	// 		AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 			},
	// 			{
	// 				ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
	// 		}},
	// 		AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 				to.Ptr("*")},
	// 				MaxAge: to.Ptr[int32](1440),
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
	// 					CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
	// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
	// 						OfferThroughput: to.Ptr[int32](1000),
	// 					},
	// 					ExportConfiguration: &armhealthcareapis.ServiceExportConfigurationInfo{
	// 						StorageAccountName: to.Ptr("existingStorageAccount"),
	// 					},
	// 					PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr(armhealthcareapis.PublicNetworkAccessDisabled),
	// 				},
	// 			}
}
