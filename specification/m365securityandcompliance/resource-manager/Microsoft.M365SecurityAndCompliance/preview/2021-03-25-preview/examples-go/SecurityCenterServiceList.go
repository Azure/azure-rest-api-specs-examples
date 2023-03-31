package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceList.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().NewListPager(nil)
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
		// page.PrivateLinkServicesForM365SecurityCenterDescriptionListResult = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescriptionListResult{
		// 	Value: []*armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		// 		{
		// 			Name: to.Ptr("service1"),
		// 			Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
		// 			Etag: to.Ptr("etag"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
		// 			Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armm365securityandcompliance.ServicesProperties{
		// 				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
		// 					{
		// 						ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
		// 					},
		// 					{
		// 						ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
		// 				}},
		// 				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
		// 					Audience: to.Ptr("https://azurehealthcareapis.com"),
		// 					Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					SmartProxyEnabled: to.Ptr(true),
		// 				},
		// 				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
		// 					AllowCredentials: to.Ptr(false),
		// 					Headers: []*string{
		// 						to.Ptr("*")},
		// 						MaxAge: to.Ptr[int64](1440),
		// 						Methods: []*string{
		// 							to.Ptr("DELETE"),
		// 							to.Ptr("GET"),
		// 							to.Ptr("OPTIONS"),
		// 							to.Ptr("PATCH"),
		// 							to.Ptr("POST"),
		// 							to.Ptr("PUT")},
		// 							Origins: []*string{
		// 								to.Ptr("*")},
		// 							},
		// 							CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
		// 								KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
		// 								OfferThroughput: to.Ptr[int64](1000),
		// 							},
		// 							PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
		// 							},
		// 							ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
		// 							PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
		// 						},
		// 				}},
		// 			}
	}
}
