package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetServiceHavingMsi.json
func ExampleServiceClient_Get_apiManagementServiceGetServiceHavingMsi() {
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
	// 	Name: to.Ptr("apimService1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Etag: to.Ptr("AAAAAAAENfI="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("ca1d33f7-0000-42ec-0000-d526a1ee953a"),
	// 		TenantID: to.Ptr("72f988bf-0000-41af-0000-2d7cd011db47"),
	// 		UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/apimService1": &armapimanagement.UserIdentityProperties{
	// 				ClientID: to.Ptr("aaff9c7d-0000-4db2-0000-ab0e3e7806cf"),
	// 				PrincipalID: to.Ptr("95194df2-9208-0000-0000-a10d2af9b5a3"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("West Europe"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-12T00:20:15.601Z"); return t}()),
	// 		CustomProperties: map[string]*string{
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("True"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("True"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("True"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("True"),
	// 			"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("True"),
	// 		},
	// 		DeveloperPortalURL: to.Ptr("https://apimService1.developer.azure-api.net"),
	// 		DisableGateway: to.Ptr(false),
	// 		GatewayRegionalURL: to.Ptr("https://apimService1-westeurope-01.regional.azure-api.net"),
	// 		GatewayURL: to.Ptr("https://apimService1.azure-api.net"),
	// 		HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 				DefaultSSLBinding: to.Ptr(false),
	// 				HostName: to.Ptr("apimService1.azure-api.net"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 			},
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 				Certificate: &armapimanagement.CertificateInformation{
	// 					Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-18T11:11:47.000Z"); return t}()),
	// 					Subject: to.Ptr("CN=*.msitesting.net"),
	// 					Thumbprint: to.Ptr("9833D531D7A45XXXXXA85908BD3692E0BD3F"),
	// 				},
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceKeyVault),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("proxy.msitesting.net"),
	// 				KeyVaultID: to.Ptr("https://samir-msi-keyvault.vault.azure.net/secrets/msicertificate"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicIPAddresses: []*string{
	// 			to.Ptr("13.94.xxx.188")},
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 			VirtualNetworkConfiguration: &armapimanagement.VirtualNetworkConfiguration{
	// 				SubnetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/dfVirtualNetwork/subnets/backendSubnet"),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeExternal),
	// 			PublisherEmail: to.Ptr("foo@contoso.com"),
	// 			PublisherName: to.Ptr("Contoso"),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 	}
}
