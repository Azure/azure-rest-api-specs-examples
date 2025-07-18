package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Accounts_GetChangeKeyVaultInformation.json
func ExampleAccountsClient_BeginGetChangeKeyVaultInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginGetChangeKeyVaultInformation(ctx, "myRG", "account1", nil)
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
	// res.GetKeyVaultStatusResponse = armnetapp.GetKeyVaultStatusResponse{
	// 	Properties: &armnetapp.GetKeyVaultStatusResponseProperties{
	// 		KeyName: to.Ptr("rsakey"),
	// 		KeyVaultPrivateEndpoints: []*armnetapp.KeyVaultPrivateEndpoint{
	// 			{
	// 				PrivateEndpointID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/privateEndpoints/privip1"),
	// 				VirtualNetworkID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 		}},
	// 		KeyVaultResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.KeyVault/managedHSMs/my-hsm"),
	// 		KeyVaultURI: to.Ptr("https://my-key-vault.managedhsm.azure.net"),
	// 	},
	// }
}
