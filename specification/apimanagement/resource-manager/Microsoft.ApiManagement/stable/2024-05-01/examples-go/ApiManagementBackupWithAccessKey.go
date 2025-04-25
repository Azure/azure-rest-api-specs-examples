package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementBackupWithAccessKey.json
func ExampleServiceClient_BeginBackup_apiManagementBackupWithAccessKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceClient().BeginBackup(ctx, "rg1", "apimService1", armapimanagement.ServiceBackupRestoreParameters{
		AccessKey:      to.Ptr("**************************************************"),
		AccessType:     to.Ptr(armapimanagement.AccessTypeAccessKey),
		BackupName:     to.Ptr("apimService1backup_2017_03_19"),
		ContainerName:  to.Ptr("backupContainer"),
		StorageAccount: to.Ptr("teststorageaccount"),
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
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXB9C2C91F1D174FDB3A2"),
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
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXB9C2C91F1D174FDB3A2"),
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
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXB9C2C91F1D174FDB3A2"),
	// 					},
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("portal1.msitesting.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeConfigurationAPI),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2036-01-01T07:00:00.000Z"); return t}()),
	// 						Subject: to.Ptr("CN=*.msitesting.net"),
	// 						Thumbprint: to.Ptr("8E989XXXXXXXXXXXXXXXXB9C2C91F1D174FDB3A2"),
	// 					},
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("configuration-api.msitesting.net"),
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
