package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/stable/2024-11-01/examples/ListCertificatesByAppServiceCertificateOrder.json
func ExampleCertificateOrdersClient_NewListCertificatesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateOrdersClient().NewListCertificatesPager("testrg123", "SampleCertificateOrderName", nil)
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
		// page.CertificateCollection = armappservice.CertificateCollection{
		// 	Value: []*armappservice.CertificateResource{
		// 		{
		// 			Name: to.Ptr("SampleCertName1"),
		// 			Type: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/certificates/SampleCertName1"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armappservice.Certificate{
		// 				KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
		// 				KeyVaultSecretName: to.Ptr("SampleSecretName1"),
		// 				ProvisioningState: to.Ptr(armappservice.KeyVaultSecretStatusSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SampleCertName2"),
		// 			Type: to.Ptr("Microsoft.CertificateRegistration/certificateOrders/certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.CertificateRegistration/certificateOrders/SampleCertificateOrderName/certificates/SampleCertName2"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armappservice.Certificate{
		// 				KeyVaultID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
		// 				KeyVaultSecretName: to.Ptr("SampleCertName2"),
		// 				ProvisioningState: to.Ptr(armappservice.KeyVaultSecretStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
