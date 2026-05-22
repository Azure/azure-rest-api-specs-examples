package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v3"
)

// Generated from example definition: 2025-04-01-preview/dicomservices/DicomServices_List.json
func ExampleDicomServicesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDicomServicesClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page = armhealthcareapis.DicomServicesClientListByWorkspaceResponse{
		// 	DicomServiceCollection: armhealthcareapis.DicomServiceCollection{
		// 		NextLink: to.Ptr("https://nextlink.url"),
		// 		Value: []*armhealthcareapis.DicomService{
		// 			{
		// 				Name: to.Ptr("blue"),
		// 				Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/blue"),
		// 				Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
		// 					Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("2f1f372f-edcf-43f5-aedb-173da3cc5c1e"),
		// 					TenantID: to.Ptr("abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 				},
		// 				Properties: &armhealthcareapis.DicomServiceProperties{
		// 					AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
		// 						Audiences: []*string{
		// 							to.Ptr("https://azurehealthcareapis.com/"),
		// 							to.Ptr("https://azurehealthcareapis.com"),
		// 						},
		// 						Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					},
		// 					CorsConfiguration: &armhealthcareapis.CorsConfiguration{
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
		// 					EnableDataPartitions: to.Ptr(false),
		// 					Encryption: &armhealthcareapis.Encryption{
		// 						CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
		// 							KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
		// 						},
		// 					},
		// 					EventState: to.Ptr(armhealthcareapis.ServiceEventStateDisabled),
		// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 					ServiceURL: to.Ptr("https://workspace1-blue.dicom.azurehealthcareapis.com"),
		// 					StorageConfiguration: &armhealthcareapis.StorageConfiguration{
		// 						FileSystemName: to.Ptr("fileSystemName"),
		// 						StorageIndexingConfiguration: &armhealthcareapis.StorageIndexingConfiguration{
		// 							StorageEventQueueName: to.Ptr("storage-events"),
		// 						},
		// 						StorageResourceID: to.Ptr("/subscriptions/ab309d4e-4c2e-4241-be2e-08e1c8dd4246/resourceGroups/rgname/providers/Microsoft.Storage/storageAccounts/accountname"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("red"),
		// 				Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/red"),
		// 				Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
		// 					Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("2f1f372f-edcf-43f5-aedb-173da3cc5c1e"),
		// 					TenantID: to.Ptr("abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 				},
		// 				Properties: &armhealthcareapis.DicomServiceProperties{
		// 					AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
		// 						Audiences: []*string{
		// 							to.Ptr("https://azurehealthcareapis.com/"),
		// 							to.Ptr("https://azurehealthcareapis.com"),
		// 						},
		// 						Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					},
		// 					EnableDataPartitions: to.Ptr(false),
		// 					Encryption: &armhealthcareapis.Encryption{
		// 						CustomerManagedKeyEncryption: &armhealthcareapis.EncryptionCustomerManagedKeyEncryption{
		// 							KeyEncryptionKeyURL: to.Ptr("https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"),
		// 						},
		// 					},
		// 					EventState: to.Ptr(armhealthcareapis.ServiceEventStateDisabled),
		// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 					ServiceURL: to.Ptr("https://workspace1-red.dicom.azurehealthcareapis.com"),
		// 					StorageConfiguration: &armhealthcareapis.StorageConfiguration{
		// 						FileSystemName: to.Ptr("fileSystemName"),
		// 						StorageIndexingConfiguration: &armhealthcareapis.StorageIndexingConfiguration{
		// 							StorageEventQueueName: to.Ptr("storage-events"),
		// 						},
		// 						StorageResourceID: to.Ptr("/subscriptions/ab309d4e-4c2e-4241-be2e-08e1c8dd4246/resourceGroups/rgname/providers/Microsoft.Storage/storageAccounts/accountname"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
