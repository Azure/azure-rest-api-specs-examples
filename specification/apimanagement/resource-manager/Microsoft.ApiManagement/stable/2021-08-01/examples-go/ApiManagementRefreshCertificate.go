package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementRefreshCertificate.json
func ExampleCertificateClient_RefreshSecret() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateClient().RefreshSecret(ctx, "rg1", "apimService1", "templateCertkv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateContract = armapimanagement.CertificateContract{
	// 	Name: to.Ptr("templateCertkv"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/certificates"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/templateCertkv"),
	// 	Properties: &armapimanagement.CertificateContractProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2037-01-01T07:00:00Z"); return t}()),
	// 		KeyVault: &armapimanagement.KeyVaultContractProperties{
	// 			IdentityClientID: to.Ptr("ceaa6b06-c00f-43ef-99ac-f53d1fe876a0"),
	// 			SecretIdentifier: to.Ptr("https://rpbvtkeyvaultintegration.vault-int.azure-int.net/secrets/msitestingCert"),
	// 			LastStatus: &armapimanagement.KeyVaultLastAccessStatusContractProperties{
	// 				Code: to.Ptr("Success"),
	// 				TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-22T00:24:53.3191468Z"); return t}()),
	// 			},
	// 		},
	// 		Subject: to.Ptr("CN=*.msitesting.net"),
	// 		Thumbprint: to.Ptr("EA**********************9AD690"),
	// 	},
	// }
}
