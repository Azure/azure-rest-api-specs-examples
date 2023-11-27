package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateMultiRegionServiceWithCustomHostname.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateMultiRegionServiceWithCustomHostname() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", armapimanagement.ServiceResource{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
			"tag3": to.Ptr("value3"),
		},
		Location: to.Ptr("West US"),
		Properties: &armapimanagement.ServiceProperties{
			AdditionalLocations: []*armapimanagement.AdditionalLocation{
				{
					DisableGateway: to.Ptr(true),
					Location:       to.Ptr("East US"),
					SKU: &armapimanagement.ServiceSKUProperties{
						Name:     to.Ptr(armapimanagement.SKUTypePremium),
						Capacity: to.Ptr[int32](1),
					},
				}},
			APIVersionConstraint: &armapimanagement.APIVersionConstraint{
				MinAPIVersion: to.Ptr("2019-01-01"),
			},
			HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
				{
					Type:                to.Ptr(armapimanagement.HostnameTypeProxy),
					CertificatePassword: to.Ptr("Password"),
					DefaultSSLBinding:   to.Ptr(true),
					EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
					HostName:            to.Ptr("gateway1.msitesting.net"),
				},
				{
					Type:                to.Ptr(armapimanagement.HostnameTypeManagement),
					CertificatePassword: to.Ptr("Password"),
					EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
					HostName:            to.Ptr("mgmt.msitesting.net"),
				},
				{
					Type:                to.Ptr(armapimanagement.HostnameTypePortal),
					CertificatePassword: to.Ptr("Password"),
					EncodedCertificate:  to.Ptr("****** Base 64 Encoded Certificate ************"),
					HostName:            to.Ptr("portal1.msitesting.net"),
				}},
			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
			PublisherEmail:     to.Ptr("apim@autorestsdk.com"),
			PublisherName:      to.Ptr("autorestsdk"),
		},
		SKU: &armapimanagement.ServiceSKUProperties{
			Name:     to.Ptr(armapimanagement.SKUTypePremium),
			Capacity: to.Ptr[int32](1),
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
	// res.ServiceResource = armapimanagement.ServiceResource{
	// 	Name: to.Ptr("apimService1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 		"tag3": to.Ptr("value3"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAACXok="),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		AdditionalLocations: []*armapimanagement.AdditionalLocation{
	// 			{
	// 				DisableGateway: to.Ptr(true),
	// 				GatewayRegionalURL: to.Ptr("https://apimService1-eastus-01.regional.azure-api.net"),
	// 				Location: to.Ptr("East US"),
	// 				PublicIPAddresses: []*string{
	// 					to.Ptr("23.101.138.153")},
	// 					SKU: &armapimanagement.ServiceSKUProperties{
	// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 						Capacity: to.Ptr[int32](1),
	// 					},
	// 			}},
	// 			APIVersionConstraint: &armapimanagement.APIVersionConstraint{
	// 				MinAPIVersion: to.Ptr("2019-01-01"),
	// 			},
	// 			CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-18T06:26:20.334Z"); return t}()),
	// 			CustomProperties: map[string]*string{
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("False"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("False"),
	// 			},
	// 			DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 			DisableGateway: to.Ptr(false),
	// 			GatewayRegionalURL: to.Ptr("https://apimService1-westus-01.regional.azure-api.net"),
	// 			GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 			HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("apimService1.azure-api.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2036-01-01T07:00:00.000Z"); return t}()),
	// 						Subject: to.Ptr("CN=*.msitesting.net"),
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXF1D174FDB3A2"),
	// 					},
	// 					DefaultSSLBinding: to.Ptr(true),
	// 					HostName: to.Ptr("gateway1.msitesting.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeManagement),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2036-01-01T07:00:00.000Z"); return t}()),
	// 						Subject: to.Ptr("CN=*.msitesting.net"),
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXF1D174FDB3A2"),
	// 					},
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("mgmt.msitesting.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypePortal),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2036-01-01T07:00:00.000Z"); return t}()),
	// 						Subject: to.Ptr("CN=*.msitesting.net"),
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXF1D174FDB3A2"),
	// 					},
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("portal1.msitesting.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 			}},
	// 			ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 			NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 			PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicIPAddresses: []*string{
	// 				to.Ptr("13.91.32.113")},
	// 				ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 				TargetProvisioningState: to.Ptr(""),
	// 				VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 				PublisherEmail: to.Ptr("apim@autorestsdk.com"),
	// 				PublisherName: to.Ptr("autorestsdk"),
	// 			},
	// 			SKU: &armapimanagement.ServiceSKUProperties{
	// 				Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 				Capacity: to.Ptr[int32](1),
	// 			},
	// 		}
}
