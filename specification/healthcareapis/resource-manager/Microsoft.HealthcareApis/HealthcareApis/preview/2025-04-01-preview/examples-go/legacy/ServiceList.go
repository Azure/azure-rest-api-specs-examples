package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v3"
)

// Generated from example definition: 2025-04-01-preview/legacy/ServiceList.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager(nil)
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
		// page = armhealthcareapis.ServicesClientListResponse{
		// 	ServicesDescriptionListResult: armhealthcareapis.ServicesDescriptionListResult{
		// 		NextLink: to.Ptr("https://host/subscriptions/subid/providers/Microsoft.HealthcareApis/services?api-version=2018-08-20-preview&%24skipToken=e30%3d"),
		// 		Value: []*armhealthcareapis.ServicesDescription{
		// 			{
		// 				Name: to.Ptr("service1"),
		// 				Type: to.Ptr("Microsoft.HealthcareApis/services"),
		// 				Etag: to.Ptr("etag"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/services/service1"),
		// 				Kind: to.Ptr(armhealthcareapis.KindFhirR4),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armhealthcareapis.ServicesProperties{
		// 					AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
		// 						{
		// 							ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
		// 						},
		// 						{
		// 							ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
		// 						},
		// 					},
		// 					AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
		// 						Audience: to.Ptr("https://azurehealthcareapis.com"),
		// 						Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 						SmartProxyEnabled: to.Ptr(true),
		// 					},
		// 					CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
		// 						AllowCredentials: to.Ptr(false),
		// 						Headers: []*string{
		// 							to.Ptr("*"),
		// 						},
		// 						MaxAge: to.Ptr[int32](1440),
		// 						Methods: []*string{
		// 							to.Ptr("DELETE"),
		// 							to.Ptr("GET"),
		// 							to.Ptr("OPTIONS"),
		// 							to.Ptr("PATCH"),
		// 							to.Ptr("POST"),
		// 							to.Ptr("PUT"),
		// 						},
		// 						Origins: []*string{
		// 							to.Ptr("*"),
		// 						},
		// 					},
		// 					CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
		// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
		// 						OfferThroughput: to.Ptr[int32](1000),
		// 					},
		// 					PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{
		// 					},
		// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateCreating),
		// 					PublicNetworkAccess: to.Ptr(armhealthcareapis.PublicNetworkAccessDisabled),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
