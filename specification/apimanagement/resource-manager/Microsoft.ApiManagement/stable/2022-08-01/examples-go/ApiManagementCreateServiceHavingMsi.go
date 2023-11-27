package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceHavingMsi.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateServiceHavingMsi() {
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
			Name:     to.Ptr(armapimanagement.SKUTypeConsumption),
			Capacity: to.Ptr[int32](0),
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
	// 	Etag: to.Ptr("AAAAAAAAWiU="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("dfb9a757-df69-4966-a8d0-711a9cd8ffb4"),
	// 		TenantID: to.Ptr("00000000-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-11T16:29:29.971Z"); return t}()),
	// 		EnableClientCertificate: to.Ptr(false),
	// 		GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 		HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("apimService1.azure-api.net"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PlatformVersion: to.Ptr(armapimanagement.PlatformVersionMtv1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TargetProvisioningState: to.Ptr(""),
	// 		VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 		PublisherEmail: to.Ptr("apim@autorestsdk.com"),
	// 		PublisherName: to.Ptr("autorestsdk"),
	// 	},
	// 	SKU: &armapimanagement.ServiceSKUProperties{
	// 		Name: to.Ptr(armapimanagement.SKUTypeConsumption),
	// 		Capacity: to.Ptr[int32](0),
	// 	},
	// 	SystemData: &armapimanagement.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 	},
	// }
}
