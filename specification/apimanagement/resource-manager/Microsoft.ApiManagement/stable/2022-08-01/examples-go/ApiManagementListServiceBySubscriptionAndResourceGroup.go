package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListServiceBySubscriptionAndResourceGroup.json
func ExampleServiceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceClient().NewListByResourceGroupPager("rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ServiceListResult = armapimanagement.ServiceListResult{
		// 	Value: []*armapimanagement.ServiceResource{
		// 		{
		// 			Name: to.Ptr("OGF-Z3-06162021-Premium"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/OGF-Z3-06162021-Premium"),
		// 			Tags: map[string]*string{
		// 				"ReleaseName": to.Ptr("Z3"),
		// 				"owner": to.Ptr("v-aswmoh"),
		// 			},
		// 			Etag: to.Ptr("AAAAAAAWN/4="),
		// 			Identity: &armapimanagement.ServiceIdentity{
		// 				Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssignedUserAssigned),
		// 				PrincipalID: to.Ptr("306205e7-b21a-41bf-92e2-3e28af30041e"),
		// 				TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 				UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
		// 					"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ogf-identity": &armapimanagement.UserIdentityProperties{
		// 						ClientID: to.Ptr("8d9791f2-0cdf-41f4-9e66-cdc39b496789"),
		// 						PrincipalID: to.Ptr("713784d2-ee37-412a-95f0-3768f397f82d"),
		// 					},
		// 				},
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armapimanagement.ServiceProperties{
		// 				AdditionalLocations: []*armapimanagement.AdditionalLocation{
		// 					{
		// 						DisableGateway: to.Ptr(false),
		// 						GatewayRegionalURL: to.Ptr("https://ogf-z3-06162021-premium-eastus2-01.regional.azure-api.net"),
		// 						Location: to.Ptr("East US 2"),
		// 						PublicIPAddresses: []*string{
		// 							to.Ptr("40.70.24.106")},
		// 							SKU: &armapimanagement.ServiceSKUProperties{
		// 								Name: to.Ptr(armapimanagement.SKUTypePremium),
		// 								Capacity: to.Ptr[int32](1),
		// 							},
		// 							Zones: []*string{
		// 							},
		// 					}},
		// 					APIVersionConstraint: &armapimanagement.APIVersionConstraint{
		// 						MinAPIVersion: to.Ptr("2019-12-01"),
		// 					},
		// 					Certificates: []*armapimanagement.CertificateConfiguration{
		// 					},
		// 					CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.945Z"); return t}()),
		// 					CustomProperties: map[string]*string{
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("false"),
		// 						"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("false"),
		// 					},
		// 					DeveloperPortalURL: to.Ptr("https://ogf-z3-06162021-premium.developer.azure-api.net"),
		// 					DisableGateway: to.Ptr(false),
		// 					GatewayRegionalURL: to.Ptr("https://ogf-z3-06162021-premium-eastus-01.regional.azure-api.net"),
		// 					GatewayURL: to.Ptr("https://ogf-z3-06162021-premium.azure-api.net"),
		// 					HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
		// 						{
		// 							Type: to.Ptr(armapimanagement.HostnameTypeProxy),
		// 							CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
		// 							DefaultSSLBinding: to.Ptr(false),
		// 							HostName: to.Ptr("ogf-z3-06162021-premium.azure-api.net"),
		// 							NegotiateClientCertificate: to.Ptr(false),
		// 						},
		// 						{
		// 							Type: to.Ptr(armapimanagement.HostnameTypeProxy),
		// 							Certificate: &armapimanagement.CertificateInformation{
		// 								Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32.000Z"); return t}()),
		// 								Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
		// 								Thumbprint: to.Ptr("BA0C286F71AF3B6A01BDB240C58A4A507E3DBD51"),
		// 							},
		// 							CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
		// 							DefaultSSLBinding: to.Ptr(true),
		// 							HostName: to.Ptr("gateway.current.int-azure-api.net"),
		// 							KeyVaultID: to.Ptr("https://ogf-testing.vault-int.azure-int.net/secrets/current-ssl"),
		// 							NegotiateClientCertificate: to.Ptr(true),
		// 						},
		// 						{
		// 							Type: to.Ptr(armapimanagement.HostnameTypeDeveloperPortal),
		// 							Certificate: &armapimanagement.CertificateInformation{
		// 								Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32.000Z"); return t}()),
		// 								Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
		// 								Thumbprint: to.Ptr("BA0C286F71AF3B6A01BDB240C58A4A507E3DBD51"),
		// 							},
		// 							CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
		// 							DefaultSSLBinding: to.Ptr(false),
		// 							HostName: to.Ptr("developer.current.int-azure-api.net"),
		// 							KeyVaultID: to.Ptr("https://ogf-testing.vault-int.azure-int.net/secrets/current-ssl"),
		// 							NegotiateClientCertificate: to.Ptr(false),
		// 						},
		// 						{
		// 							Type: to.Ptr(armapimanagement.HostnameTypeManagement),
		// 							Certificate: &armapimanagement.CertificateInformation{
		// 								Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-08T22:32:32.000Z"); return t}()),
		// 								Subject: to.Ptr("CN=*.current.int-azure-api.net, O=Microsoft Corporation, L=Redmond, S=WA, C=US"),
		// 								Thumbprint: to.Ptr("BA0C286F71AF3B6A01BDB240C58A4A507E3DBD51"),
		// 							},
		// 							CertificateSource: to.Ptr(armapimanagement.CertificateSourceCustom),
		// 							DefaultSSLBinding: to.Ptr(false),
		// 							HostName: to.Ptr("mgmt.current.int-azure-api.net"),
		// 							KeyVaultID: to.Ptr("https://ogf-testing.vault-int.azure-int.net/secrets/current-ssl"),
		// 							NegotiateClientCertificate: to.Ptr(false),
		// 					}},
		// 					ManagementAPIURL: to.Ptr("https://ogf-z3-06162021-premium.management.azure-api.net"),
		// 					NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
		// 					PortalURL: to.Ptr("https://ogf-z3-06162021-premium.portal.azure-api.net"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					PublicIPAddresses: []*string{
		// 						to.Ptr("13.92.130.49")},
		// 						PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
		// 						ScmURL: to.Ptr("https://ogf-z3-06162021-premium.scm.azure-api.net"),
		// 						TargetProvisioningState: to.Ptr(""),
		// 						VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
		// 						PublisherEmail: to.Ptr("bar@contoso.com"),
		// 						PublisherName: to.Ptr("Test Premium"),
		// 					},
		// 					SKU: &armapimanagement.ServiceSKUProperties{
		// 						Name: to.Ptr(armapimanagement.SKUTypePremium),
		// 						Capacity: to.Ptr[int32](1),
		// 					},
		// 					SystemData: &armapimanagement.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:40:00.710Z"); return t}()),
		// 						CreatedBy: to.Ptr("bar@contoso.com"),
		// 						CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T06:33:09.615Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("foo@contoso.com"),
		// 						LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("vvktestcons"),
		// 					Type: to.Ptr("Microsoft.ApiManagement/service"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/vvktestcons"),
		// 					Tags: map[string]*string{
		// 						"Owner": to.Ptr("vitaliik"),
		// 					},
		// 					Etag: to.Ptr("AAAAAAAWKwo="),
		// 					Location: to.Ptr("East US"),
		// 					Properties: &armapimanagement.ServiceProperties{
		// 						CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-10T18:07:23.456Z"); return t}()),
		// 						CustomProperties: map[string]*string{
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("False"),
		// 						},
		// 						DisableGateway: to.Ptr(false),
		// 						EnableClientCertificate: to.Ptr(false),
		// 						GatewayURL: to.Ptr("https://vvktestcons.azure-api.net"),
		// 						HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
		// 							{
		// 								Type: to.Ptr(armapimanagement.HostnameTypeProxy),
		// 								CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
		// 								DefaultSSLBinding: to.Ptr(true),
		// 								HostName: to.Ptr("vvktestcons.azure-api.net"),
		// 								NegotiateClientCertificate: to.Ptr(false),
		// 						}},
		// 						NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
		// 						TargetProvisioningState: to.Ptr(""),
		// 						VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
		// 						PublisherEmail: to.Ptr("bar@contoso.com"),
		// 						PublisherName: to.Ptr("vvktestcons"),
		// 					},
		// 					SKU: &armapimanagement.ServiceSKUProperties{
		// 						Name: to.Ptr(armapimanagement.SKUTypeConsumption),
		// 						Capacity: to.Ptr[int32](0),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("OGF-Z3-06162021-Standard"),
		// 					Type: to.Ptr("Microsoft.ApiManagement/service"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/OGF-Z3-06162021-Standard"),
		// 					Tags: map[string]*string{
		// 					},
		// 					Etag: to.Ptr("AAAAAAAWF7M="),
		// 					Identity: &armapimanagement.ServiceIdentity{
		// 						Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssignedUserAssigned),
		// 						PrincipalID: to.Ptr("347a5800-ca99-475a-9202-fe38ca79ee41"),
		// 						TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 						UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
		// 							"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ogf-identity": &armapimanagement.UserIdentityProperties{
		// 								ClientID: to.Ptr("8d9791f2-0cdf-41f4-9e66-cdc39b496789"),
		// 								PrincipalID: to.Ptr("713784d2-ee37-412a-95f0-3768f397f82d"),
		// 							},
		// 						},
		// 					},
		// 					Location: to.Ptr("East US"),
		// 					Properties: &armapimanagement.ServiceProperties{
		// 						APIVersionConstraint: &armapimanagement.APIVersionConstraint{
		// 							MinAPIVersion: to.Ptr("2019-12-01"),
		// 						},
		// 						CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:39:58.655Z"); return t}()),
		// 						CustomProperties: map[string]*string{
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("false"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA": to.Ptr("true"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("False"),
		// 							"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("False"),
		// 						},
		// 						DeveloperPortalURL: to.Ptr("https://ogf-z3-06162021-standard.developer.azure-api.net"),
		// 						DisableGateway: to.Ptr(false),
		// 						GatewayRegionalURL: to.Ptr("https://ogf-z3-06162021-standard-eastus-01.regional.azure-api.net"),
		// 						GatewayURL: to.Ptr("https://ogf-z3-06162021-standard.azure-api.net"),
		// 						HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
		// 							{
		// 								Type: to.Ptr(armapimanagement.HostnameTypeProxy),
		// 								CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
		// 								DefaultSSLBinding: to.Ptr(true),
		// 								HostName: to.Ptr("ogf-z3-06162021-standard.azure-api.net"),
		// 								NegotiateClientCertificate: to.Ptr(false),
		// 						}},
		// 						ManagementAPIURL: to.Ptr("https://ogf-z3-06162021-standard.management.azure-api.net"),
		// 						NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
		// 						PortalURL: to.Ptr("https://ogf-z3-06162021-standard.portal.azure-api.net"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						PublicIPAddresses: []*string{
		// 							to.Ptr("13.82.208.32")},
		// 							PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
		// 							ScmURL: to.Ptr("https://ogf-z3-06162021-standard.scm.azure-api.net"),
		// 							TargetProvisioningState: to.Ptr(""),
		// 							VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
		// 							PublisherEmail: to.Ptr("bar@contoso.com"),
		// 							PublisherName: to.Ptr("Test Standard"),
		// 						},
		// 						SKU: &armapimanagement.ServiceSKUProperties{
		// 							Name: to.Ptr(armapimanagement.SKUTypeStandard),
		// 							Capacity: to.Ptr[int32](2),
		// 						},
		// 						SystemData: &armapimanagement.SystemData{
		// 							CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T09:39:58.459Z"); return t}()),
		// 							CreatedBy: to.Ptr("bar@contoso.com"),
		// 							CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 							LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T15:05:13.549Z"); return t}()),
		// 							LastModifiedBy: to.Ptr("bar@contoso.com"),
		// 							LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("ogf-dev-060921"),
		// 						Type: to.Ptr("Microsoft.ApiManagement/service"),
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/ogf-dev-060921"),
		// 						Tags: map[string]*string{
		// 						},
		// 						Etag: to.Ptr("AAAAAAAWEFg="),
		// 						Identity: &armapimanagement.ServiceIdentity{
		// 							Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssigned),
		// 							PrincipalID: to.Ptr("c9bd4c05-205e-4431-b232-112cf2e9e0aa"),
		// 							TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
		// 						},
		// 						Location: to.Ptr("East US"),
		// 						Properties: &armapimanagement.ServiceProperties{
		// 							AdditionalLocations: []*armapimanagement.AdditionalLocation{
		// 								{
		// 									DisableGateway: to.Ptr(false),
		// 									GatewayRegionalURL: to.Ptr("https://ogf-dev-060921-southcentralus-01.regional.azure-api.net"),
		// 									Location: to.Ptr("South Central US"),
		// 									PublicIPAddresses: []*string{
		// 										to.Ptr("13.84.208.29")},
		// 										SKU: &armapimanagement.ServiceSKUProperties{
		// 											Name: to.Ptr(armapimanagement.SKUTypePremium),
		// 											Capacity: to.Ptr[int32](9),
		// 										},
		// 										Zones: []*string{
		// 										},
		// 								}},
		// 								APIVersionConstraint: &armapimanagement.APIVersionConstraint{
		// 									MinAPIVersion: to.Ptr("2019-12-01"),
		// 								},
		// 								Certificates: []*armapimanagement.CertificateConfiguration{
		// 									{
		// 										Certificate: &armapimanagement.CertificateInformation{
		// 											Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T22:51:47.000Z"); return t}()),
		// 											Subject: to.Ptr("CN=*.apim.net"),
		// 											Thumbprint: to.Ptr("4E8234312EC69245D1AE296C4882D46FB84076A3"),
		// 										},
		// 										StoreName: to.Ptr(armapimanagement.CertificateConfigurationStoreNameRoot),
		// 								}},
		// 								CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T10:06:22.213Z"); return t}()),
		// 								CustomProperties: map[string]*string{
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Ssl30": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Ssl30": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10": to.Ptr("false"),
		// 									"Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11": to.Ptr("false"),
		// 								},
		// 								DeveloperPortalURL: to.Ptr("https://ogf-dev-060921.developer.azure-api.net"),
		// 								DisableGateway: to.Ptr(false),
		// 								GatewayRegionalURL: to.Ptr("https://ogf-dev-060921-eastus-01.regional.azure-api.net"),
		// 								GatewayURL: to.Ptr("https://ogf-dev-060921.azure-api.net"),
		// 								HostnameConfigurations: []*armapimanagement.HostnameConfiguration{
		// 									{
		// 										Type: to.Ptr(armapimanagement.HostnameTypeProxy),
		// 										CertificateSource: to.Ptr(armapimanagement.CertificateSourceBuiltIn),
		// 										DefaultSSLBinding: to.Ptr(true),
		// 										HostName: to.Ptr("ogf-dev-060921.azure-api.net"),
		// 										NegotiateClientCertificate: to.Ptr(false),
		// 								}},
		// 								ManagementAPIURL: to.Ptr("https://ogf-dev-060921.management.azure-api.net"),
		// 								NotificationSenderEmail: to.Ptr("apimgmt-noreply@mail.windowsazure.com"),
		// 								PortalURL: to.Ptr("https://ogf-dev-060921.portal.azure-api.net"),
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 								PublicIPAddresses: []*string{
		// 									to.Ptr("168.62.39.172")},
		// 									PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
		// 									ScmURL: to.Ptr("https://ogf-dev-060921.scm.azure-api.net"),
		// 									TargetProvisioningState: to.Ptr(""),
		// 									VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
		// 									PublisherEmail: to.Ptr("v-ssaiprasan@microsoft.com"),
		// 									PublisherName: to.Ptr("TechM"),
		// 								},
		// 								SKU: &armapimanagement.ServiceSKUProperties{
		// 									Name: to.Ptr(armapimanagement.SKUTypePremium),
		// 									Capacity: to.Ptr[int32](3),
		// 								},
		// 								SystemData: &armapimanagement.SystemData{
		// 									CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T10:06:21.733Z"); return t}()),
		// 									CreatedBy: to.Ptr("v-ssaiprasan@microsoft.com"),
		// 									CreatedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 									LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-10T14:27:05.546Z"); return t}()),
		// 									LastModifiedBy: to.Ptr("v-ssaiprasan@microsoft.com"),
		// 									LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
		// 								},
		// 						}},
		// 					}
	}
}
