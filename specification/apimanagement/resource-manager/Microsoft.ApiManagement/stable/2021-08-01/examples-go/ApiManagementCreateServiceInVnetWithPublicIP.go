package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateServiceInVnetWithPublicIP.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateServiceInVnetWithPublicIp() {
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
		Location: to.Ptr("East US 2 EUAP"),
		Properties: &armapimanagement.ServiceProperties{
			PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/publicIPAddresses/apimazvnet"),
			VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
				SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/virtualNetworks/apimcus/subnets/tenant"),
			},
			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeExternal),
			PublisherEmail:     to.Ptr("apim@autorestsdk.com"),
			PublisherName:      to.Ptr("autorestsdk"),
		},
		SKU: &armapimanagement.ServiceSKUProperties{
			Name:     to.Ptr(armapimanagement.SKUTypePremium),
			Capacity: to.Ptr[int32](2),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2")},
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
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.ApiManagement/service/apimService1"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 		"tag3": to.Ptr("value3"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAAGTAs="),
	// 	Location: to.Ptr("East US 2 EUAP"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-22T06:53:46.6409875Z"); return t}()),
	// 		CustomProperties: map[string]*string{
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("False"),
	// 		},
	// 		DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 		DisableGateway: to.Ptr(false),
	// 		GatewayRegionalURL: to.Ptr("https://apimService1-eastus2euap-01.regional.azure-api.net"),
	// 		GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 		HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("apimService1.azure-api.net"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 		PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicIPAddresses: []*string{
	// 			to.Ptr("20.47.137.XXX")},
	// 			PublicIPAddressID: to.Ptr("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/publicIPAddresses/apimazvnet"),
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 			VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 				SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rgName/providers/Microsoft.Network/virtualNetworks/apimcus/subnets/tenant"),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeExternal),
	// 			PublisherEmail: to.Ptr("apim@autorestsdk.com"),
	// 			PublisherName: to.Ptr("autorestsdk"),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 			Capacity: to.Ptr[int32](2),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2")},
	// 		}
}
