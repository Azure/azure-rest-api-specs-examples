package armdomainservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateDomainService.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestResourceGroup",
		"TestDomainService.com",
		armdomainservices.DomainService{
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
