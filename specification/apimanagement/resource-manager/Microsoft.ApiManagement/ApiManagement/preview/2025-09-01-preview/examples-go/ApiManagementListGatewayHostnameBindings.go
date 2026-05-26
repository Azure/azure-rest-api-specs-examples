package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListGatewayHostnameBindings.json
func ExampleAPIGatewayHostnameBindingClient_NewListByGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIGatewayHostnameBindingClient().NewListByGatewayPager("rg1", "standard-gw-1", nil)
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
		// page = armapimanagement.APIGatewayHostnameBindingClientListByGatewayResponse{
		// 	GatewayHostnameBindingListResult: armapimanagement.GatewayHostnameBindingListResult{
		// 		Value: []*armapimanagement.GatewayHostnameBindingResource{
		// 			{
		// 				Name: to.Ptr("hb-1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/gateways/hostnameBindings"),
		// 				ETag: to.Ptr("AAAAAAAWN/4="),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1/hostnameBindings/hb-1"),
		// 				Properties: &armapimanagement.GatewayHostnameBindingBaseProperties{
		// 					Certificate: &armapimanagement.GatewayHostnameBindingCertificate{
		// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 						Subject: to.Ptr("CN=primary.cotoso.com"),
		// 						Thumbprint: to.Ptr("3F2504E0-4F89-11D3-9A0C-0305E82C3301"),
		// 					},
		// 					Hostname: to.Ptr("primary.cotoso.com"),
		// 					KeyVault: &armapimanagement.GatewayHostnameBindingKeyVault{
		// 						IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						LastStatus: &armapimanagement.GatewayHostnameBindingKeyVaultLastStatus{
		// 							Code: to.Ptr(armapimanagement.KeyVaultFetchCodeSuccess),
		// 							TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-02T12:00:00Z"); return t}()),
		// 						},
		// 						SecretID: to.Ptr("https://myvault.keyvault.azure.net/secrets/contosoprimarycert"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("hb-2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/gateways/hostnameBindings"),
		// 				ETag: to.Ptr("AAAAAAAWKwo="),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/gateways/standard-gw-1/hostnameBindings/hb-2"),
		// 				Properties: &armapimanagement.GatewayHostnameBindingBaseProperties{
		// 					Certificate: &armapimanagement.GatewayHostnameBindingCertificate{
		// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
		// 						Subject: to.Ptr("CN=backup.cotoso.com"),
		// 						Thumbprint: to.Ptr("3F2504E0-4F89-11D3-9A0C-0305E82C3301"),
		// 					},
		// 					Hostname: to.Ptr("backup.cotoso.com"),
		// 					KeyVault: &armapimanagement.GatewayHostnameBindingKeyVault{
		// 						IdentityClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						LastStatus: &armapimanagement.GatewayHostnameBindingKeyVaultLastStatus{
		// 							Code: to.Ptr(armapimanagement.KeyVaultFetchCodeFailed),
		// 							LastSuccessTimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T12:00:00Z"); return t}()),
		// 							TimeStampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-02T12:00:00Z"); return t}()),
		// 						},
		// 						SecretID: to.Ptr("https://myvault.keyvault.azure.net/secrets/contosobackupcert"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
