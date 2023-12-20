package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/fhirservices/FhirServices_Create.json
func ExampleFhirServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFhirServicesClient().BeginCreateOrUpdate(ctx, "testRG", "workspace1", "fhirservice1", armhealthcareapis.FhirService{
		Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
			Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armhealthcareapis.UserAssignedIdentity{
				"/subscriptions/subid/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-mi": {},
			},
		},
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"additionalProp1": to.Ptr("string"),
			"additionalProp2": to.Ptr("string"),
			"additionalProp3": to.Ptr("string"),
		},
		Kind: to.Ptr(armhealthcareapis.FhirServiceKindFhirR4),
		Properties: &armhealthcareapis.FhirServiceProperties{
			AcrConfiguration: &armhealthcareapis.FhirServiceAcrConfiguration{
				LoginServers: []*string{
					to.Ptr("test1.azurecr.io")},
			},
			AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
				Audience:          to.Ptr("https://azurehealthcareapis.com"),
				Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
				SmartProxyEnabled: to.Ptr(true),
			},
			CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
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
			Encryption: &armhealthcareapis.Encryption{
				CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
					KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
				},
			},
			ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
				StorageAccountName: to.Ptr("existingStorageAccount"),
			},
			ImplementationGuidesConfiguration: &armhealthcareapis.ImplementationGuidesConfiguration{
				UsCoreMissingData: to.Ptr(false),
			},
			ImportConfiguration: &armhealthcareapis.FhirServiceImportConfiguration{
				Enabled:              to.Ptr(false),
				InitialImportMode:    to.Ptr(false),
				IntegrationDataStore: to.Ptr("existingStorageAccount"),
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
	// res.FhirService = armhealthcareapis.FhirService{
	// 	Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
	// 		Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armhealthcareapis.UserAssignedIdentity{
	// 			"/subscriptions/subid/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-mi": &armhealthcareapis.UserAssignedIdentity{
	// 			},
	// 		},
	// 	},
	// 	Name: to.Ptr("fhirservice1"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirservices/fhirservice1"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"additionalProp1": to.Ptr("string"),
	// 		"additionalProp2": to.Ptr("string"),
	// 		"additionalProp3": to.Ptr("string"),
	// 	},
	// 	Kind: to.Ptr(armhealthcareapis.FhirServiceKindFhirR4),
	// 	Properties: &armhealthcareapis.FhirServiceProperties{
	// 		AuthenticationConfiguration: &armhealthcareapis.FhirServiceAuthenticationConfiguration{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armhealthcareapis.FhirServiceCorsConfiguration{
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
	// 					Encryption: &armhealthcareapis.Encryption{
	// 						CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
	// 							KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
	// 						},
	// 					},
	// 					EventState: to.Ptr(armhealthcareapis.ServiceEventStateDisabled),
	// 					ExportConfiguration: &armhealthcareapis.FhirServiceExportConfiguration{
	// 						StorageAccountName: to.Ptr("existingStorageAccount"),
	// 					},
	// 					ImplementationGuidesConfiguration: &armhealthcareapis.ImplementationGuidesConfiguration{
	// 						UsCoreMissingData: to.Ptr(false),
	// 					},
	// 					ImportConfiguration: &armhealthcareapis.FhirServiceImportConfiguration{
	// 						Enabled: to.Ptr(false),
	// 						InitialImportMode: to.Ptr(false),
	// 						IntegrationDataStore: to.Ptr("existingStorageAccount"),
	// 					},
	// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 				},
	// 			}
}
