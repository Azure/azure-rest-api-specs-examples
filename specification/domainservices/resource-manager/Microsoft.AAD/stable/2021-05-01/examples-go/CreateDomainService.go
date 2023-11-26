package armdomainservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateDomainService.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "TestResourceGroup", "TestDomainService.com", armdomainservices.DomainService{
		Properties: &armdomainservices.DomainServiceProperties{
			DomainName: to.Ptr("TestDomainService.com"),
			DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
				NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
				SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
				TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
			},
			FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
			LdapsSettings: &armdomainservices.LdapsSettings{
				ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
				Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
				PfxCertificate:         to.Ptr("MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w..."),
				PfxCertificatePassword: to.Ptr("<pfxCertificatePassword>"),
			},
			NotificationSettings: &armdomainservices.NotificationSettings{
				AdditionalRecipients: []*string{
					to.Ptr("jicha@microsoft.com"),
					to.Ptr("caalmont@microsoft.com")},
				NotifyDcAdmins:     to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
				NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
			},
			ReplicaSets: []*armdomainservices.ReplicaSet{
				{
					Location: to.Ptr("West US"),
					SubnetID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS"),
				}},
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
	// res.DomainService = armdomainservices.DomainService{
	// 	Name: to.Ptr("TestDomainService.com"),
	// 	Type: to.Ptr("Microsoft.AAD/DomainServices"),
	// 	ID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestResourceGroup/providers/Microsoft.AAD/DomainServices/TestDomainService.com"),
	// 	Properties: &armdomainservices.DomainServiceProperties{
	// 		DeploymentID: to.Ptr("4a619871-0150-41c4-aeb4-0b10deb7940a"),
	// 		DomainName: to.Ptr("TestDomainService.com"),
	// 		DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
	// 			NtlmV1: to.Ptr(armdomainservices.NtlmV1Enabled),
	// 			SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
	// 			TLSV1: to.Ptr(armdomainservices.TLSV1Disabled),
	// 		},
	// 		FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
	// 		LdapsSettings: &armdomainservices.LdapsSettings{
	// 			CertificateNotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-15T21:43:21.000Z"); return t}()),
	// 			CertificateThumbprint: to.Ptr("9154A390F0C387D679E0DD040701745CDFED67F3"),
	// 			ExternalAccess: to.Ptr(armdomainservices.ExternalAccessEnabled),
	// 			Ldaps: to.Ptr(armdomainservices.LdapsEnabled),
	// 			PublicCertificate: to.Ptr("MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w..."),
	// 		},
	// 		NotificationSettings: &armdomainservices.NotificationSettings{
	// 			AdditionalRecipients: []*string{
	// 				to.Ptr("jicha@microsoft.com"),
	// 				to.Ptr("caalmont@microsoft.com")},
	// 				NotifyDcAdmins: to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
	// 				NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ReplicaSets: []*armdomainservices.ReplicaSet{
	// 				{
	// 					DomainControllerIPAddress: []*string{
	// 						to.Ptr("10.0.0.1"),
	// 						to.Ptr("10.0.0.2")},
	// 						ExternalAccessIPAddress: to.Ptr("13.64.148.151"),
	// 						Location: to.Ptr("West US"),
	// 						ReplicaSetID: to.Ptr("4a619871-0150-41c4-aeb4-0b10deb7940a"),
	// 						ServiceStatus: to.Ptr("Running"),
	// 						SubnetID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS"),
	// 						VnetSiteID: to.Ptr("99083198-a39c-469f-972d-59017e7f078c"),
	// 				}},
	// 				SyncOwner: to.Ptr("4a619871-0150-41c4-aeb4-0b10deb7940a"),
	// 				TenantID: to.Ptr("3f8cd22c-7b32-48aa-a01c-f533133b1def"),
	// 				Version: to.Ptr[int32](2),
	// 			},
	// 		}
}
