package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRCustomCertificates_List.json
func ExampleCustomCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomCertificatesClient().NewListPager("myResourceGroup", "mySignalRService", nil)
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
		// page.CustomCertificateList = armsignalr.CustomCertificateList{
		// 	Value: []*armsignalr.CustomCertificate{
		// 		{
		// 			Name: to.Ptr("myCert"),
		// 			Type: to.Ptr("Microsoft.SignalRService/SignalR/customCertificates"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/customCertificates/myCert"),
		// 			Properties: &armsignalr.CustomCertificateProperties{
		// 				KeyVaultBaseURI: to.Ptr("https://myvault.keyvault.azure.net/"),
		// 				KeyVaultSecretName: to.Ptr("mycert"),
		// 				KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
		// 				ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
