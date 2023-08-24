package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateService.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateService() {
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
			"Name": to.Ptr("Contoso"),
			"Test": to.Ptr("User"),
		},
		Location: to.Ptr("South Central US"),
		Properties: &armapimanagement.ServiceProperties{
			PublisherEmail: to.Ptr("foo@contoso.com"),
			PublisherName:  to.Ptr("foo"),
		},
		SKU: &armapimanagement.ServiceSKUProperties{
			Name:     to.Ptr(armapimanagement.SKUTypeDeveloper),
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
	// 		"api-version": to.Ptr("2022-08-01"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAAmREI="),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:01.2506031Z"); return t}()),
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
	// 		GatewayRegionalURL: to.Ptr("https://apimService1-eastus-01.regional.azure-api.net"),
	// 		GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 		HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("apimService1.azure-api.net"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv1),
	// 		PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicIPAddresses: []*string{
	// 			to.Ptr("13.90.229.33")},
	// 			PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 			PublisherEmail: to.Ptr("samir@microsoft.com"),
	// 			PublisherName: to.Ptr("foo"),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypeStandard),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.9390609Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-11T18:41:00.9390609Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		},
	// 	}
}
