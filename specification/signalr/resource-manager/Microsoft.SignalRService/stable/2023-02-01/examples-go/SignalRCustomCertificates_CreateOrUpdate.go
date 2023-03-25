package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalRCustomCertificates_CreateOrUpdate.json
func ExampleCustomCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomCertificatesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "mySignalRService", "myCert", armsignalr.CustomCertificate{
		Properties: &armsignalr.CustomCertificateProperties{
			KeyVaultBaseURI:       to.Ptr("https://myvault.keyvault.azure.net/"),
			KeyVaultSecretName:    to.Ptr("mycert"),
			KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
		},
	}, nil)
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
