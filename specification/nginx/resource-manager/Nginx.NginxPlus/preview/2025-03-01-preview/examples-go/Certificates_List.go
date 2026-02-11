package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/Certificates_List.json
func ExampleCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificatesClient().NewListPager("myResourceGroup", "myDeployment", nil)
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
		// page = armnginx.CertificatesClientListResponse{
		// 	CertificateListResponse: armnginx.CertificateListResponse{
		// 		Value: []*armnginx.Certificate{
		// 			{
		// 				Name: to.Ptr("cert1"),
		// 				Type: to.Ptr("nginx.nginxplus/nginxdeployments/certificates"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/certificates/cert1"),
		// 				Properties: &armnginx.CertificateProperties{
		// 					CertificateError: &armnginx.CertificateErrorResponseBody{
		// 						Code: to.Ptr("CertificateNotYetValid"),
		// 						Message: to.Ptr("The certificate is not yet valid. The certificate is not yet valid. NotBefore: '2020-01-01T00:00:00Z', NotAfter: '2020-01-01T00:00:00Z'"),
		// 					},
		// 					CertificateVirtualPath: to.Ptr("/src/cert/somePath.cert"),
		// 					KeyVaultSecretCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
		// 					KeyVaultSecretID: to.Ptr("https://someKV.vault.azure.com/someSecretID"),
		// 					KeyVaultSecretVersion: to.Ptr("12345678ef9a12345678ef9a12345678"),
		// 					KeyVirtualPath: to.Ptr("/src/cert/somekey.key"),
		// 					ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 					SHA1Thumbprint: to.Ptr("1234567890ABCDEF1234567890ABCDEF12345678"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("cert2"),
		// 				Type: to.Ptr("nginx.nginxplus/nginxdeployments/certificates"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/certificates/cert2"),
		// 				Properties: &armnginx.CertificateProperties{
		// 					CertificateVirtualPath: to.Ptr("/src/cert/somePath2.cert"),
		// 					KeyVaultSecretCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
		// 					KeyVaultSecretID: to.Ptr("https://someKV.vault.azure.com/someSecretID2"),
		// 					KeyVaultSecretVersion: to.Ptr("12345678ef9a12345678ef9a12345678"),
		// 					KeyVirtualPath: to.Ptr("/src/cert/somekey2.key"),
		// 					ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 					SHA1Thumbprint: to.Ptr("1234567890ABCDEF1234567890ABCDEF12345678"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
