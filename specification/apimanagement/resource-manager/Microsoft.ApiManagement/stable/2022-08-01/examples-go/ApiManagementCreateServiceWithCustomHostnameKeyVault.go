package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateServiceWithCustomHostnameKeyVault.json
func ExampleServiceClient_BeginCreateOrUpdate_apiManagementCreateServiceWithCustomHostnameKeyVault() {
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
			Type: to.Ptr(armapimanagement.ApimIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
				"/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Location: to.Ptr("North Europe"),
		Properties: &armapimanagement.ServiceProperties{
			APIVersionConstraint: &armapimanagement.APIVersionConstraint{
				MinAPIVersion: to.Ptr("2019-01-01"),
			},
			HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
				{
					Type:              to.Ptr(armapimanagement.HostnameTypeProxy),
					DefaultSSLBinding: to.Ptr(true),
					HostName:          to.Ptr("gateway1.msitesting.net"),
					IdentityClientID:  to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultID:        to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
				},
				{
					Type:             to.Ptr(armapimanagement.HostnameTypeManagement),
					HostName:         to.Ptr("mgmt.msitesting.net"),
					IdentityClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultID:       to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
				},
				{
					Type:             to.Ptr(armapimanagement.HostnameTypePortal),
					HostName:         to.Ptr("portal1.msitesting.net"),
					IdentityClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
					KeyVaultID:       to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
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
	// 	Etag: to.Ptr("AAAAAAAigjU="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeUserAssigned),
	// 		TenantID: to.Ptr("f686d426-8d16-0000-0000-ab578e110ccd"),
	// 		UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armapimanagement.UserIdentityProperties{
	// 				ClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
	// 				PrincipalID: to.Ptr("15e769b2-0000-0000-0000-3fd9a923ac3a"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("North Europe"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		APIVersionConstraint: &armapimanagement.APIVersionConstraint{
	// 			MinAPIVersion: to.Ptr("2019-01-01"),
	// 		},
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-13T22:30:20.7759747Z"); return t}()),
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
	// 		GatewayRegionalURL: to.Ptr("https://apimService1-northeurope-01.regional.azure-api.net"),
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
	// 					Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2037-01-01T07:00:00+00:00"); return t}()),
	// 					Subject: to.Ptr("CN=*.msitesting.net"),
	// 					Thumbprint: to.Ptr("EA276907917CB5XXXXXXXXXXX690"),
	// 				},
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceKeyVault),
	// 				DefaultSSLBinding: to.Ptr(true),
	// 				HostName: to.Ptr("gateway1.msitesting.net"),
	// 				IdentityClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
	// 				KeyVaultID: to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 			},
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypeManagement),
	// 				Certificate: &armapimanagement.CertificateInformation{
	// 					Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2037-01-01T07:00:00+00:00"); return t}()),
	// 					Subject: to.Ptr("CN=*.msitesting.net"),
	// 					Thumbprint: to.Ptr("EA276907917CB5XXXXXXXXXXX690"),
	// 				},
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceKeyVault),
	// 				DefaultSSLBinding: to.Ptr(false),
	// 				HostName: to.Ptr("mgmt.msitesting.net"),
	// 				IdentityClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
	// 				KeyVaultID: to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 			},
	// 			{
	// 				Type: to.Ptr(armapimanagement.HostnameTypePortal),
	// 				Certificate: &armapimanagement.CertificateInformation{
	// 					Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2037-01-01T07:00:00+00:00"); return t}()),
	// 					Subject: to.Ptr("CN=*.msitesting.net"),
	// 					Thumbprint: to.Ptr("EA276907917CB5XXXXXXXXXXX690"),
	// 				},
	// 				CertificateSource: to.Ptr(armapimanagement.CertificateSourceKeyVault),
	// 				DefaultSSLBinding: to.Ptr(false),
	// 				HostName: to.Ptr("portal1.msitesting.net"),
	// 				IdentityClientID: to.Ptr("329419bc-adec-4dce-9568-25a6d486e468"),
	// 				KeyVaultID: to.Ptr("https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert"),
	// 				NegotiateClientCertificate: to.Ptr(false),
	// 		}},
	// 		ManagementAPIURL: to.Ptr("https://apimService1.management.azure-api.net"),
	// 		NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 		PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 		PortalURL: to.Ptr("https://apimService1.portal.azure-api.net"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PublicIPAddresses: []*string{
	// 			to.Ptr("40.112.74.192")},
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 			PublisherEmail: to.Ptr("apim@autorestsdk.com"),
	// 			PublisherName: to.Ptr("autorestsdk"),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-02T02:03:01.1974346Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeApplication),
	// 		},
	// 	}
}
