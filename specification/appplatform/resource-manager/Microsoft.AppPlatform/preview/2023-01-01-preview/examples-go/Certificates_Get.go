package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/Certificates_Get.json
func ExampleCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Get(ctx, "myResourceGroup", "myservice", "mycertificate", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateResource = armappplatform.CertificateResource{
	// 	Name: to.Ptr("mycertificate"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/certificates"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycertificate"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.KeyVaultCertificateProperties{
	// 		Type: to.Ptr("KeyVaultCertificate"),
	// 		ActivateDate: to.Ptr("2019-02-22T07:40:42Z"),
	// 		DNSNames: []*string{
	// 			to.Ptr("mydomain.com"),
	// 			to.Ptr("mydomain.net"),
	// 			to.Ptr("mydomain.io")},
	// 			ExpirationDate: to.Ptr("2019-02-21T07:40:42Z"),
	// 			IssuedDate: to.Ptr("2019-02-20T07:40:42Z"),
	// 			Issuer: to.Ptr("mydomain.com"),
	// 			ProvisioningState: to.Ptr(armappplatform.CertificateResourceProvisioningStateSucceeded),
	// 			SubjectName: to.Ptr("mysubjectname"),
	// 			Thumbprint: to.Ptr("934367bf1c97033f877db0f15cb1b586957d3133"),
	// 			CertVersion: to.Ptr("08a219d06d874795a96db47e06fbb01e"),
	// 			ExcludePrivateKey: to.Ptr(false),
	// 			KeyVaultCertName: to.Ptr("mycert"),
	// 			VaultURI: to.Ptr("https://myvault.vault.azure.net"),
	// 		},
	// 	}
}
