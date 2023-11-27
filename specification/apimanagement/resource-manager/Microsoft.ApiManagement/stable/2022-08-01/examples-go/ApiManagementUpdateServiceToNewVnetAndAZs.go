package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateServiceToNewVnetAndAZs.json
func ExampleServiceClient_BeginUpdate_apiManagementUpdateServiceToNewVnetAndAvailabilityZones() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginUpdate(ctx, "rg1", "apimService1", armapimanagement.ServiceUpdateParameters{
		Properties: &armapimanagement.ServiceUpdateProperties{
			AdditionalLocations: []*armapimanagement.AdditionalLocation{
				{
					Location:          to.Ptr("Australia East"),
					PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/apim-australia-east-publicip"),
					SKU: &armapimanagement.ServiceSKUProperties{
						Name:     to.Ptr(armapimanagement.SKUTypePremium),
						Capacity: to.Ptr[int32](3),
					},
					VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
						SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimaeavnet/subnets/default"),
					},
					Zones: []*string{
						to.Ptr("1"),
						to.Ptr("2"),
						to.Ptr("3")},
				}},
			PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/publicip-apim-japan-east"),
			VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
				SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-apim-japaneast/subnets/apim2"),
			},
			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeExternal),
		},
		SKU: &armapimanagement.ServiceSKUProperties{
			Name:     to.Ptr(armapimanagement.SKUTypePremium),
			Capacity: to.Ptr[int32](3),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2"),
			to.Ptr("3")},
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
	// res.ServiceResource = armapimanagement.ServiceResource{
	// 	Name: to.Ptr("apimService1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 	Etag: to.Ptr("AAAAAAAWBIU="),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		AdditionalLocations: []*armapimanagement.AdditionalLocation{
	// 			{
	// 				DisableGateway: to.Ptr(false),
	// 				GatewayRegionalURL: to.Ptr("https://apimService1-australiaeast-01.regional.azure-api.net"),
	// 				Location: to.Ptr("Australia East"),
	// 				PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 				PublicIPAddresses: []*string{
	// 					to.Ptr("20.213.1.35")},
	// 					PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/apim-australia-east-publicip"),
	// 					SKU: &armapimanagement.ServiceSKUProperties{
	// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 						Capacity: to.Ptr[int32](3),
	// 					},
	// 					VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 						SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apimaeavnet/subnets/default"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("1"),
	// 						to.Ptr("2"),
	// 						to.Ptr("3")},
	// 				}},
	// 				CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-08T23:41:35.644Z"); return t}()),
	// 				CustomProperties: map[string]*string{
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_GCM_SHA256": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_256_CBC_SHA": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_256_CBC_SHA256": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("false"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("false"),
	// 				},
	// 				DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 				DisableGateway: to.Ptr(false),
	// 				GatewayRegionalURL: to.Ptr("https://apimService1-japaneast-01.regional.azure-api.net"),
	// 				GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 				HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 					{
	// 						Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 						CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 						DefaultSSLBinding: to.Ptr(false),
	// 						HostName: to.Ptr("apimService1.azure-api.net"),
	// 						NegotiateClientCertificate: to.Ptr(false),
	// 					},
	// 					{
	// 						Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 						Certificate: &armapimanagement.CertificateInformation{
	// 							Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-09T23:59:59.000Z"); return t}()),
	// 							Subject: to.Ptr("CN=mycustomdomain.int-azure-api.net"),
	// 							Thumbprint: to.Ptr("2994B5FFB8F76B3C687D324A8DEE0432C1ED18CD"),
	// 						},
	// 						CertificateSource: to.Ptr(armapimanagement.CertificateSourceManaged),
	// 						DefaultSSLBinding: to.Ptr(true),
	// 						HostName: to.Ptr("mycustomdomain.int-azure-api.net"),
	// 						NegotiateClientCertificate: to.Ptr(false),
	// 				}},
	// 				ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 				NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 				PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 				PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				PublicIPAddresses: []*string{
	// 					to.Ptr("20.78.248.217")},
	// 					PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/publicip-apim-japan-east"),
	// 					PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 					ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 					TargetProvisioningState: to.Ptr(""),
	// 					VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 						SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-apim-japaneast/subnets/apim2"),
	// 					},
	// 					VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeInternal),
	// 					PublisherEmail: to.Ptr("contoso@microsoft.com"),
	// 					PublisherName: to.Ptr("apimPublisher"),
	// 				},
	// 				SKU: &armapimanagement.ServiceSKUProperties{
	// 					Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 					Capacity: to.Ptr[int32](3),
	// 				},
	// 				SystemData: &armapimanagement.SystemData{
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-21T20:04:21.610Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("contoso@microsoft.com"),
	// 					LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 				},
	// 				Zones: []*string{
	// 					to.Ptr("1"),
	// 					to.Ptr("2"),
	// 					to.Ptr("3")},
	// 				}
}
