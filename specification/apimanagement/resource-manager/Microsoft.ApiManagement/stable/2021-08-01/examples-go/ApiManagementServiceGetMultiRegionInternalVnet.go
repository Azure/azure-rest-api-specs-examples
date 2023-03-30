package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetMultiRegionInternalVnet.json
func ExampleServiceClient_Get_apiManagementServiceGetMultiRegionInternalVnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().Get(ctx, "rg1", "apimService1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armapimanagement.ServiceResource{
	// 	Name: to.Ptr("apimservice1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimservice1"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("AAAAAADqC0c="),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		AdditionalLocations: []*armapimanagement.AdditionalLocation{
	// 			{
	// 				DisableGateway: to.Ptr(false),
	// 				GatewayRegionalURL: to.Ptr("https://apimservice1-westus2-01.regional.azure-api.net"),
	// 				Location: to.Ptr("West US 2"),
	// 				PrivateIPAddresses: []*string{
	// 					to.Ptr("10.0.X.6")},
	// 					PublicIPAddresses: []*string{
	// 						to.Ptr("40.XXX.79.187")},
	// 						SKU: &armapimanagement.ServiceSKUProperties{
	// 							Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 							Capacity: to.Ptr[int32](1),
	// 						},
	// 						VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 							SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/APIMVNet/subnets/apim-internal-sub"),
	// 						},
	// 				}},
	// 				APIVersionConstraint: &armapimanagement.APIVersionConstraint{
	// 				},
	// 				CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-02T01:42:09.1268424Z"); return t}()),
	// 				CustomProperties: map[string]*string{
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("True"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("True"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("True"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("True"),
	// 					"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("True"),
	// 				},
	// 				DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 				DisableGateway: to.Ptr(false),
	// 				GatewayRegionalURL: to.Ptr("https://apimservice1-westus-01.regional.azure-api.net"),
	// 				GatewayURL: to.Ptr("https://apimservice1.azure-api.net"),
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
	// 							Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-16T16:51:34+00:00"); return t}()),
	// 							Subject: to.Ptr("CN=*.preview.net"),
	// 							Thumbprint: to.Ptr("B4330123DBAXXXXXXXXX1F35E84493476"),
	// 						},
	// 						CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
	// 						DefaultSSLBinding: to.Ptr(true),
	// 						HostName: to.Ptr("apimgatewaytest.preview.net"),
	// 						NegotiateClientCertificate: to.Ptr(false),
	// 				}},
	// 				ManagementAPIURL: to.Ptr("https://apimservice1.management.azure-api.net"),
	// 				NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 				PortalURL: to.Ptr("https://apimservice1.portal.azure-api.net"),
	// 				PrivateIPAddresses: []*string{
	// 					to.Ptr("172.XX.0.5")},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					PublicIPAddresses: []*string{
	// 						to.Ptr("137.XXX.11.74")},
	// 						ScmURL: to.Ptr("https://apimservice1.scm.azure-api.net"),
	// 						TargetProvisioningState: to.Ptr(""),
	// 						VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 							SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/apim-appGateway-vnet/subnets/apim-subnet"),
	// 						},
	// 						VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeInternal),
	// 						PublisherEmail: to.Ptr("abcs@contoso.com"),
	// 						PublisherName: to.Ptr("contoso publisher"),
	// 					},
	// 					SKU: &armapimanagement.ServiceSKUProperties{
	// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 						Capacity: to.Ptr[int32](1),
	// 					},
	// 				}
}
