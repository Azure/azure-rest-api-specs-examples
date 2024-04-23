package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/fhirservices/FhirServices_Get.json
func ExampleFhirServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFhirServicesClient().Get(ctx, "testRG", "workspace1", "fhirservices1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FhirService = armhealthcareapis.FhirService{
	// 	Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
	// 		Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("2f1f372f-edcf-43f5-aedb-173da3cc5c1e"),
	// 		TenantID: to.Ptr("abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 	},
	// 	Name: to.Ptr("fhirservices1"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirservices/fhirservices1"),
	// 	Properties: &armhealthcareapis.FhirServiceProperties{
	// 		AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartIdentityProviders: []*armhealthcareapis.SmartIdentityProviderConfiguration{
	// 				{
	// 					Applications: []*armhealthcareapis.SmartIdentityProviderApplication{
	// 						{
	// 							AllowedDataActions: []*armhealthcareapis.SmartDataActions{
	// 								to.Ptr(armhealthcareapis.SmartDataActionsRead)},
	// 								Audience: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 								ClientID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 						}},
	// 						Authority: to.Ptr("https://login.b2clogin.com/11111111-1111-1111-1111-111111111111/v2.0"),
	// 				}},
	// 				SmartProxyEnabled: to.Ptr(true),
	// 			},
	// 			CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
	// 				AllowCredentials: to.Ptr(false),
	// 				Headers: []*string{
	// 					to.Ptr("*")},
	// 					MaxAge: to.Ptr[int32](1440),
	// 					Methods: []*string{
	// 						to.Ptr("DELETE"),
	// 						to.Ptr("GET"),
	// 						to.Ptr("OPTIONS"),
	// 						to.Ptr("PATCH"),
	// 						to.Ptr("POST"),
	// 						to.Ptr("PUT")},
	// 						Origins: []*string{
	// 							to.Ptr("*")},
	// 						},
	// 						Encryption: &armhealthcareapis.Encryption{
	// 							CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
	// 								KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
	// 							},
	// 						},
	// 						ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
	// 							StorageAccountName: to.Ptr("existingStorageAccount"),
	// 						},
	// 						ImplementationGuidesConfiguration: &armhealthcareapis.ImplementationGuidesConfiguration{
	// 							UsCoreMissingData: to.Ptr(false),
	// 						},
	// 						ImportConfiguration: &armhealthcareapis.FhirServiceImportConfiguration{
	// 							Enabled: to.Ptr(false),
	// 							InitialImportMode: to.Ptr(false),
	// 							IntegrationDataStore: to.Ptr("existingStorageAccount"),
	// 						},
	// 						ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 					},
	// 				}
}
