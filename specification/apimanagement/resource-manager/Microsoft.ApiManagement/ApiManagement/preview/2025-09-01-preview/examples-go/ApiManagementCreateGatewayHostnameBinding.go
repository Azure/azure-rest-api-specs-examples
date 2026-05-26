package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateGatewayHostnameBinding.json
func ExampleAPIGatewayHostnameBindingClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIGatewayHostnameBindingClient().BeginCreateOrUpdate(ctx, "rg1", "apimGateway1", "hb-1", armapimanagement.GatewayHostnameBindingResource{
		Properties: &armapimanagement.GatewayHostnameBindingBaseProperties{
			Hostname: to.Ptr("primary.cotoso.com"),
			KeyVault: &armapimanagement.GatewayHostnameBindingKeyVault{
				IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
				SecretID:         to.Ptr("https://myvault.keyvault.azure.net/secrets/contosoprimarycert"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIGatewayHostnameBindingClientCreateOrUpdateResponse{
	// 	GatewayHostnameBindingResource: armapimanagement.GatewayHostnameBindingResource{
	// 		Name: to.Ptr("hb-1"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/gateways/hostnameBindings"),
	// 		ETag: to.Ptr("AAAAAAAmREI"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1/hostnameBindings/hb-1"),
	// 		Properties: &armapimanagement.GatewayHostnameBindingBaseProperties{
	// 			Certificate: &armapimanagement.GatewayHostnameBindingCertificate{
	// 				Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
	// 				Subject: to.Ptr("CN=primary.cotoso.com"),
	// 				Thumbprint: to.Ptr("3F2504E0-4F89-11D3-9A0C-0305E82C3301"),
	// 			},
	// 			Hostname: to.Ptr("primary.cotoso.com"),
	// 			KeyVault: &armapimanagement.GatewayHostnameBindingKeyVault{
	// 				IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				LastStatus: &armapimanagement.GatewayHostnameBindingKeyVaultLastStatus{
	// 					Code: to.Ptr(armapimanagement.KeyVaultFetchCodeSuccess),
	// 					TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-02T12:00:00Z"); return t}()),
	// 				},
	// 				SecretID: to.Ptr("https://myvault.keyvault.azure.net/secrets/contosoprimarycert"),
	// 			},
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 		},
	// 	},
	// }
}
