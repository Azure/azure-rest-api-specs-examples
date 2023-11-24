package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2023-04-01/examples/Certificates_List.json
func ExampleCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.CertificateListResponse = armnginx.CertificateListResponse{
		// 	Value: []*armnginx.Certificate{
		// 		{
		// 			Name: to.Ptr("cert1"),
		// 			Type: to.Ptr("nginx.nginxplus/nginxdeployments/certificates"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/NGINX.NGINXPLUS/nginxDeployments/myDeployment/certificates/cert1"),
		// 			Properties: &armnginx.CertificateProperties{
		// 				CertificateVirtualPath: to.Ptr("/src/cert/somePath.cert"),
		// 				KeyVaultSecretID: to.Ptr("https://someKV.vault.azure.com/someSecretID"),
		// 				KeyVirtualPath: to.Ptr("/src/cert/somekey.key"),
		// 				ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("cert2"),
		// 			Type: to.Ptr("nginx.nginxplus/nginxdeployments/certificates"),
		// 			ID: to.Ptr("/subscritions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/NGINX.NGINXPLUS/nginxDeployments/myDeployment/certificates/cert2"),
		// 			Properties: &armnginx.CertificateProperties{
		// 				CertificateVirtualPath: to.Ptr("/src/cert/somePath2.cert"),
		// 				KeyVaultSecretID: to.Ptr("https://someKV.vault.azure.com/someSecretID2"),
		// 				KeyVirtualPath: to.Ptr("/src/cert/somekey2.key"),
		// 				ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
