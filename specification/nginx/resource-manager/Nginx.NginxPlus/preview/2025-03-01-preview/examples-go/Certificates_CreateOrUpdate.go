package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/Certificates_CreateOrUpdate.json
func ExampleCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCertificatesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDeployment", "default", armnginx.Certificate{}, nil)
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
	// res = armnginx.CertificatesClientCreateOrUpdateResponse{
	// 	Certificate: &armnginx.Certificate{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("nginx.nginxplus/nginxdeployments/certificates"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/certificates/default"),
	// 		Properties: &armnginx.CertificateProperties{
	// 			CertificateVirtualPath: to.Ptr("/src/cert/somePath.cert"),
	// 			KeyVaultSecretID: to.Ptr("https://someKV.vault.azure.com/someSecretID"),
	// 			KeyVirtualPath: to.Ptr("/src/cert/somekey.key"),
	// 			ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
