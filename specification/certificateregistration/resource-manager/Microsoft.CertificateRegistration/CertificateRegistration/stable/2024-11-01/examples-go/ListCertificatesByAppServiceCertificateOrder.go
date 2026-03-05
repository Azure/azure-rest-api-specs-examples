package armcertificateregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/certificateregistration/armcertificateregistration"
)

// Generated from example definition: 2024-11-01/ListCertificatesByAppServiceCertificateOrder.json
func ExampleAppServiceCertificateOrdersClient_NewListCertificatesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcertificateregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppServiceCertificateOrdersClient().NewListCertificatesPager("testrg123", "SampleCertificateOrderName", nil)
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
		// page = armcertificateregistration.AppServiceCertificateOrdersClientListCertificatesResponse{
		// 	AppServiceCertificateCollection: armcertificateregistration.AppServiceCertificateCollection{
		// 		Value: []*armcertificateregistration.AppServiceCertificateResource{
		// 			{
		// 				Name: to.Ptr("SampleCertName1"),
		// 				Type: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/certificates/SampleCertName1"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armcertificateregistration.AppServiceCertificate{
		// 					KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
		// 					KeyVaultSecretName: to.Ptr("SampleSecretName1"),
		// 					ProvisioningState: to.Ptr(armcertificateregistration.KeyVaultSecretStatusSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("SampleCertName2"),
		// 				Type: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/certificates/SampleCertName2"),
		// 				Location: to.Ptr("Global"),
		// 				Properties: &armcertificateregistration.AppServiceCertificate{
		// 					KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
		// 					KeyVaultSecretName: to.Ptr("SampleCertName2"),
		// 					ProvisioningState: to.Ptr(armcertificateregistration.KeyVaultSecretStatusKeyVaultSecretDoesNotExist),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
