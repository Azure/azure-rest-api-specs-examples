package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementServiceRefreshKeyVaultHostnames.json
func ExampleServiceClient_BeginRefreshHostnames() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginRefreshHostnames(ctx, "rg1", "apimService1", nil)
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
	// res = armapimanagement.ServiceClientRefreshHostnamesResponse{
	// 	ServiceResource: armapimanagement.ServiceResource{
	// 		Name: to.Ptr("apimService1"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 		Etag: to.Ptr(azcore.ETag("AAAAAAK5GpM=")),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 		Identity: &armapimanagement.ServiceIdentity{
	// 			Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("b413456e-a9c8-4242-9754-0a01c061bd41"),
	// 			TenantID: to.Ptr("b413456e-a9c8-4242-9754-0a01c061bd41"),
	// 		},
	// 		Location: to.Ptr("West Europe"),
	// 		Properties: &armapimanagement.ServiceProperties{
	// 			AdditionalLocations: []*armapimanagement.AdditionalLocation{
	// 				{
	// 					DisableGateway: to.Ptr(false),
	// 					GatewayRegionalURL: to.Ptr("https://apimService1-northeurope-01.regional.azure-api.net"),
	// 					Location: to.Ptr("North Europe"),
	// 					OutboundPublicIPAddresses: []*string{
	// 						to.Ptr("4.xxx.40.176"),
	// 					},
	// 					PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 					PublicIPAddresses: []*string{
	// 						to.Ptr("4.xxx.40.176"),
	// 					},
	// 					SKU: &armapimanagement.ServiceSKUProperties{
	// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 						Capacity: to.Ptr[int32](1),
	// 					},
	// 					VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 						SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimVnet/subnets/default2"),
	// 					},
	// 				},
	// 			},
	// 			CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-12T00:20:15.6018952Z"); return t}()),
	// 			CustomProperties: map[string]*string{
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("True"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("True"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("True"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("True"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("True"),
	// 			},
	// 			DeveloperPortalStatus: to.Ptr(armapimanagement.DeveloperPortalStatusEnabled),
	// 			DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 			DisableGateway: to.Ptr(false),
	// 			GatewayRegionalURL: to.Ptr("https://apimService1-westeurope-01.regional.azure-api.net"),
	// 			GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 			HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("apimService1.azure-api.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-18T11:11:47+00:00"); return t}()),
	// 						Subject: to.Ptr("CN=*.msitesting.net"),
	// 						Thumbprint: to.Ptr("9833D531D7A45XXXXXXXXXXXX8BD392E0BD3F"),
	// 					},
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceKeyVault),
	// 					DefaultSSLBinding: to.Ptr(true),
	// 					HostName: to.Ptr("proxy.msitesting.net"),
	// 					KeyVaultID: to.Ptr("https://apim-msi-keyvault.vault.azure.net/secrets/sslcertificate"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 			},
	// 			LegacyPortalStatus: to.Ptr(armapimanagement.LegacyPortalStatusDisabled),
	// 			ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 			NatGatewayState: to.Ptr(armapimanagement.NatGatewayStateDisabled),
	// 			NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 			OutboundPublicIPAddresses: []*string{
	// 				to.Ptr("57.xxx.61.xx"),
	// 			},
	// 			PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv21),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicIPAddresses: []*string{
	// 				to.Ptr("XX.153.XX.94"),
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 			PublisherEmail: to.Ptr("autorest@contoso.com"),
	// 			PublisherName: to.Ptr("Microsoft"),
	// 			ReleaseChannel: to.Ptr(armapimanagement.ReleaseChannelDefault),
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 				SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1/subnets/default"),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeExternal),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-06-14T17:33:34.8195595Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("autorest@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
