package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackupWithUserAssignedManagedIdentity.json
func ExampleServiceClient_BeginBackup_apiManagementBackupWithUserAssignedManagedIdentity() {
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
		AccessType:     to.Ptr(armapimanagement.AccessTypeUserAssignedManagedIdentity),
		BackupName:     to.Ptr("backup5"),
		ClientID:       to.Ptr("XXXXX-a154-4830-XXXX-46a12da1a1e2"),
		ContainerName:  to.Ptr("apim-backups"),
		StorageAccount: to.Ptr("contosorpstorage"),
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
	// 		"Owner": to.Ptr("apimService1"),
	// 	},
	// 	Etag: to.Ptr("AAAAAAAQM8o="),
	// 	Identity: &armapimanagement.ServiceIdentity{
	// 		Type: to.Ptr(armapimanagement.ApimIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("00000000-5fb4-4916-95d4-64b306f9d924"),
	// 		TenantID: to.Ptr("00000000-86f1-0000-91ab-2d7cd011db47"),
	// 		UserAssignedIdentities: map[string]*armapimanagement.UserIdentityProperties{
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/rg1UserIdentity": &armapimanagement.UserIdentityProperties{
	// 				ClientID: to.Ptr("00000000-a154-4830-0000-46a12da1a1e2"),
	// 				PrincipalID: to.Ptr("00000000-a100-4478-0000-d65d98118ba0"),
	// 			},
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/vpnpremium": &armapimanagement.UserIdentityProperties{
	// 				ClientID: to.Ptr("00000000-6328-4db2-0000-ab0e3e7806cf"),
	// 				PrincipalID: to.Ptr("00000000-9208-4128-af2d-a10d2af9b5a3"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("Central US EUAP"),
	// 	Properties: &armapimanagement.ServiceProperties{
	// 		CreatedAtUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-23T16:26:47.8637967Z"); return t}()),
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
	// 		GatewayRegionalURL: to.Ptr("https://apimService1-centraluseuap-01.regional.azure-api.net"),
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
	// 			to.Ptr("52.XXXX.160.66")},
	// 			PublicNetworkAccess: to.Ptr(armapimanagement.PublicNetworkAccessEnabled),
	// 			ScmURL: to.Ptr("https://apimService1.scm.azure-api.net"),
	// 			TargetProvisioningState: to.Ptr(""),
	// 			VirtualNetworkType: to.Ptr(armapimanagement.VirtualNetworkTypeNone),
	// 			PublisherEmail: to.Ptr("apimService1@corp.microsoft.com"),
	// 			PublisherName: to.Ptr("MS"),
	// 		},
	// 		SKU: &armapimanagement.ServiceSKUProperties{
	// 			Name: to.Ptr(armapimanagement.SKUTypePremium),
	// 			Capacity: to.Ptr[int32](1),
	// 		},
	// 		SystemData: &armapimanagement.SystemData{
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-30T06:24:57.0008037Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("contoso@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armapimanagement.CreatedByTypeUser),
	// 		},
	// 	}
}
