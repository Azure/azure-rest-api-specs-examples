package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListCertificates.json
func ExampleCertificateClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificateClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.CertificateClientListByServiceOptions{Filter: nil,
		Top:                     nil,
		Skip:                    nil,
		IsKeyVaultRefreshFailed: nil,
	})
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
		// page.CertificateCollection = armapimanagement.CertificateCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.CertificateContract{
		// 		{
		// 			Name: to.Ptr("templateCert1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/certificates"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/kjoshiarmtemplateCert1"),
		// 			Properties: &armapimanagement.CertificateContractProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-23T17:03:41.000Z"); return t}()),
		// 				Subject: to.Ptr("CN=mutual-authcert"),
		// 				Thumbprint: to.Ptr("EBA************************48594A6"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("templateCertkv"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/certificates"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/templateCertkv"),
		// 			Properties: &armapimanagement.CertificateContractProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2037-01-01T07:00:00.000Z"); return t}()),
		// 				KeyVault: &armapimanagement.KeyVaultContractProperties{
		// 					IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
		// 					SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert"),
		// 					LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
		// 						Code: to.Ptr("Success"),
		// 						TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-22T00:24:53.319Z"); return t}()),
		// 					},
		// 				},
		// 				Subject: to.Ptr("CN=*.msitesting.net"),
		// 				Thumbprint: to.Ptr("EA**********************9AD690"),
		// 			},
		// 	}},
		// }
	}
}
