package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountCertificates_Get.json
func ExampleIntegrationAccountCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountCertificatesClient().Get(ctx, "testResourceGroup", "testIntegrationAccount", "testCertificate", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountCertificate = armlogic.IntegrationAccountCertificate{
	// 	Name: to.Ptr("testCertificate"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/IntegrationAccount9472/certificates/testCertificate"),
	// 	Properties: &armlogic.IntegrationAccountCertificateProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T20:42:21.051Z"); return t}()),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T20:42:21.051Z"); return t}()),
	// 		Key: &armlogic.KeyVaultKeyReference{
	// 			KeyName: to.Ptr("<keyName>"),
	// 			KeyVault: &armlogic.KeyVaultKeyReferenceKeyVault{
	// 				Name: to.Ptr("<keyVaultName>"),
	// 				Type: to.Ptr("Microsoft.KeyVault/vaults"),
	// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testResourceGroup/providers/microsoft.keyvault/vaults/<keyVaultName>"),
	// 			},
	// 			KeyVersion: to.Ptr("87d9764197604449b9b8eb7bd8710868"),
	// 		},
	// 		PublicCertificate: to.Ptr("<publicCertificate>"),
	// 	},
	// }
}
