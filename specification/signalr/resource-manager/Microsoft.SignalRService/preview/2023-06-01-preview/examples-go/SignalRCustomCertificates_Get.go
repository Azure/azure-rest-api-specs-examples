package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRCustomCertificates_Get.json
func ExampleCustomCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomCertificatesClient().Get(ctx, "myResourceGroup", "mySignalRService", "myCert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomCertificate = armsignalr.CustomCertificate{
	// 	Name: to.Ptr("myCert"),
	// 	Type: to.Ptr("Microsoft.SignalRService/SignalR/customCertificates"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/customCertificates/myCert"),
	// 	Properties: &armsignalr.CustomCertificateProperties{
	// 		KeyVaultBaseURI: to.Ptr("https://myvault.keyvault.azure.net/"),
	// 		KeyVaultSecretName: to.Ptr("mycert"),
	// 		KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
	// 		ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
	// 	},
	// }
}
