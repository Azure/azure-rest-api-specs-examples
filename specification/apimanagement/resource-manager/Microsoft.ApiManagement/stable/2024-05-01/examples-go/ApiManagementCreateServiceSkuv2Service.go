package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateServiceSkuv2Service.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateServiceSkuv2Service() {
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
		Identity: &armapimanagement.ServiceIdentity{
			Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssigned),
		},
		Location: to.Ptr("West US"),
		Properties: &armapimanagement.ServiceProperties{
			PublisherEmail: to.Ptr("apim@autorestsdk.com"),
			PublisherName:  to.Ptr("autorestsdk"),
		},
		SKU: &armapimanagement.ServiceSKUProperties{
			Name:     to.Ptr(armapimanagement.SKUTypeStandardV2),
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
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("AAAAAAA3fHM="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("de161222-0000-0000-0000-1caa5d9f0b0e"),
	// 		TenantID: to.Ptr("72f988bf-0000-0000-0000-2d7cd011db47"),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-10T18:50:52.550Z"); return t}()),
	// 		CustomProperties: map[string]*string{
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("False"),
	// 		},
	// 		DisableGateway: to.Ptr(false),
	// 		GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 		HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("apimService1.azure-api.net"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		NatGatewayState: to.Ptr(armapimanagement.NatGatewayState("Unsupported")),
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PlatformVersion: to.Ptr(armapimanagement.PlatformVersionUndetermined),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 		TargetProvisioningState: to.Ptr(""),
	// 		VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 		PublisherEmail: to.Ptr("apim@contoso.com"),
	// 		PublisherName: to.Ptr("apimgmt-skuv2"),
	// 	},
	// 	SKU: &armapimanagement.ServiceSKUProperties{
	// 		Name: to.Ptr(armapimanagement.SKUTypeStandardV2),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// 	SystemData: &armapimanagement.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-10T18:50:51.539Z"); return t}()),
	// 		CreatedBy: to.Ptr("contoso@microsoft.com"),
	// 		CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-10T18:50:51.539Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("contoso@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 	},
	// }
}
