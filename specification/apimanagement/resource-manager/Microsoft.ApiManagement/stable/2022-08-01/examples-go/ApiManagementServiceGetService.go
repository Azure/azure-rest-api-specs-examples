package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetService.json
func ExampleServiceClient_Get_apiManagementServiceGetService() {
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
	// 	Name: to.Ptr("OGF-Z3-06162021-Premium"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/OGF-Z3-06162021-Premium"),
	// 	Tags: map[string]*string{
	// 		"ReleaseName": to.Ptr("Z3"),
	// 		"owner": to.Ptr("v-aswmoh"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAAWN/4="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("306205e7-b21a-41bf-92e2-3e28af30041e"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 		UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ogf-identity": &armapimanagement.UserIdentityProperties{
	// 				ClientID: to.Ptr("8d9791f2-0cdf-41f4-9e66-cdc39b496789"),
	// 				PrincipalID: to.Ptr("713784d2-ee37-412a-95f0-3768f397f82d"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		AdditionalLocations: []*armapimanagement.AdditionalLocation{
	// 			{
	// 				DisableGateway: to.Ptr(false),
	// 				GatewayRegionalURL: to.Ptr("https://ogf-z3-06162021-premium-eastus2-01.regional.azure-api.net"),
	// 				Location: to.Ptr("East US 2"),
	// 				PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 				PublicIPAddresses: []*string{
	// 					to.Ptr("40.70.24.106")},
	// 					SKU: &armapimanagement.ServiceSKUProperties{
	// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 						Capacity: to.Ptr[int32](1),
	// 					},
	// 					Zones: []*string{
	// 					},
	// 			}},
	// 			APIVersionConstraint: &armapimanagement.APIVersionConstraint{
	// 				MinAPIVersion: to.Ptr("2019-12-01"),
	// 			},
	// 			Certificates: []*armapimanagement.CertificateConfiguration{
	// 			},
	// 			CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.9453556Z"); return t}()),
	// 			CustomProperties: map[string]*string{
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("false"),
	// 				"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("false"),
	// 			},
	// 			DeveloperPortalURL: to.Ptr("https://ogf-z3-06162021-premium.developer.azure-api.net"),
	// 			DisableGateway: to.Ptr(false),
	// 			GatewayRegionalURL: to.Ptr("https://ogf-z3-06162021-premium-eastus-01.regional.azure-api.net"),
	// 			GatewayURL: to.Ptr("https://ogf-z3-06162021-premium.azure-api.net"),
	// 			HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("ogf-z3-06162021-premium.azure-api.net"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeProxy),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32+00:00"); return t}()),
	// 						Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
	// 						Thumbprint: to.Ptr("BA0C286XXXXXXXX58A4A507E3DBD51"),
	// 					},
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
	// 					DefaultSSLBinding: to.Ptr(true),
	// 					HostName: to.Ptr("gateway.current.int-azure-api.net"),
	// 					KeyVaultID: to.Ptr("https://ogf-testing.vault.azure.net/secrets/current-ssl"),
	// 					NegotiateClientCertificate: to.Ptr(true),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeDeveloperPortal),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32+00:00"); return t}()),
	// 						Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
	// 						Thumbprint: to.Ptr("BA0C286XXXXXXXX58A4A507E3DBD51"),
	// 					},
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("developer.current.int-azure-api.net"),
	// 					KeyVaultID: to.Ptr("https://ogf-testing.vault.azure.net/secrets/current-ssl"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 				},
	// 				{
	// 					Type: to.Ptr(armapimanagement.HostnameTypeManagement),
	// 					Certificate: &armapimanagement.CertificateInformation{
	// 						Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32+00:00"); return t}()),
	// 						Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
	// 						Thumbprint: to.Ptr("BA0C286XXXXXXXX58A4A507E3DBD51"),
	// 					},
	// 					CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
	// 					DefaultSSLBinding: to.Ptr(false),
	// 					HostName: to.Ptr("mgmt.current.int-azure-api.net"),
	// 					KeyVaultID: to.Ptr("https://ogf-testing.vault.azure.net/secrets/current-ssl"),
	// 					NegotiateClientCertificate: to.Ptr(false),
	// 			}},
	// 			ManagementAPIURL: to.Ptr("https://ogf-z3-06162021-premium.management.azure-api.net"),
	// 			NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
	// 			PlatformVersion: to.Ptr(armapimanagement.PlatformVersionStv2),
	// 			PortalURL: to.Ptr("https://ogf-z3-06162021-premium.portal.azure-api.net"),
	// 			PrivateEndpointConnections: []*armapimanagement.RemotePrivateEndpointConnectionWrapper{
	// 				{
	// 					Name: to.Ptr("privateEndpointProxyName"),
	// 					Type: to.Ptr("Microsoft.ApiManagement/service/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/privateEndpointProxyName"),
	// 					Properties: &armapimanagement.PrivateEndpointConnectionWrapperProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("Gateway")},
	// 							PrivateEndpoint: &armapimanagement.ArmIDWrapper{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
	// 							},
	// 							PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
	// 								Description: to.Ptr("Please approve my request, thanks"),
	// 								ActionsRequired: to.Ptr("None"),
	// 								Status: to.Ptr(armapimanagement.PrivateEndpointServiceConnectionStatusPending),
	// 							},
	// 							ProvisioningState: to.Ptr("Succeeded"),
	// 						},
	// 				}},
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				PublicIPAddresses: []*string{
	// 					to.Ptr("13.92.130.49")},
	// 					PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 					ScmURL: to.Ptr("https://ogf-z3-06162021-premium.scm.azure-api.net"),
	// 					TargetProvisioningState: to.Ptr(""),
	// 					VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 					PublisherEmail: to.Ptr("string"),
	// 					PublisherName: to.Ptr("Test Premium"),
	// 				},
	// 				SKU: &armapimanagement.ServiceSKUProperties{
	// 					Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 					Capacity: to.Ptr[int32](1),
	// 				},
	// 				SystemData: &armapimanagement.SystemData{
	// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.7106733Z"); return t}()),
	// 					CreatedBy: to.Ptr("string"),
	// 					CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T06:33:09.6159006Z"); return t}()),
	// 					LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 					LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 				},
	// 			}
}
