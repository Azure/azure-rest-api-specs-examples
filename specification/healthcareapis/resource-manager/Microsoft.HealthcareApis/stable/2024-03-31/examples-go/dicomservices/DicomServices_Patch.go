package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/dicomservices/DicomServices_Patch.json
func ExampleDicomServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDicomServicesClient().BeginUpdate(ctx, "testRG", "blue", "workspace1", armhealthcareapis.DicomServicePatchResource{
		Tags: map[string]*string{
			"tagKey": to.Ptr("tagValue"),
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
	// res.DicomService = armhealthcareapis.DicomService{
	// 	Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
	// 		Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("2f1f372f-edcf-43f5-aedb-173da3cc5c1e"),
	// 		TenantID: to.Ptr("abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 	},
	// 	Name: to.Ptr("blue"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/blue"),
	// 	Tags: map[string]*string{
	// 		"tagKey": to.Ptr("tagValue"),
	// 	},
	// 	Properties: &armhealthcareapis.DicomServiceProperties{
	// 		AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
	// 			Audiences: []*string{
	// 				to.Ptr("https://azurehealthcareapis.com/"),
	// 				to.Ptr("https://azurehealthcareapis.com")},
	// 				Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			},
	// 			CorsConfiguration: &armhealthcareapis.CorsConfiguration{
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
	// 						EnableDataPartitions: to.Ptr(false),
	// 						Encryption: &armhealthcareapis.Encryption{
	// 							CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
	// 								KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
	// 							},
	// 						},
	// 						EventState: to.Ptr(armhealthcareapis.ServiceEventStateDisabled),
	// 						ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 						ServiceURL: to.Ptr("https://workspace1-blue.dicom.azurehealthcareapis.com"),
	// 						StorageConfiguration: &armhealthcareapis.StorageConfiguration{
	// 							FileSystemName: to.Ptr("fileSystemName"),
	// 							StorageResourceID: to.Ptr("/subscriptions/ab309d4e-4c2e-4241-be2e-08e1c8dd4246/resourceGroups/rgname/providers/Microsoft.Storage/storageAccounts/accountname"),
	// 						},
	// 					},
	// 				}
}
