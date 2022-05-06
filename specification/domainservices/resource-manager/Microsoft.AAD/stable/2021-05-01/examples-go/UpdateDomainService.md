Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdomainservices%2Farmdomainservices%2Fv0.4.0/sdk/resourcemanager/domainservices/armdomainservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armdomainservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				ConfigDiagnostics: &armdomainservices.ConfigDiagnostics{
					LastExecuted: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2021-05-05T12:00:23Z;"); return t }()),
					ValidatorResults: []*armdomainservices.ConfigDiagnosticsValidatorResult{
						{
							Issues: []*armdomainservices.ConfigDiagnosticsValidatorResultIssue{
								{
									DescriptionParams: []*string{},
									ID:                to.Ptr("<id>"),
								}},
							ReplicaSetSubnetDisplayName: to.Ptr("<replica-set-subnet-display-name>"),
							Status:                      to.Ptr(armdomainservices.StatusWarning),
							ValidatorID:                 to.Ptr("<validator-id>"),
						}},
				},
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
					SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
					TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
				},
				FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
					Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
					PfxCertificate:         to.Ptr("<pfx-certificate>"),
					PfxCertificatePassword: to.Ptr("<pfx-certificate-password>"),
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
						Location: to.Ptr("<location>"),
						SubnetID: to.Ptr("<subnet-id>"),
					},
					{
						Location: to.Ptr("<location>"),
						SubnetID: to.Ptr("<subnet-id>"),
					}},
			},
		},
		&armdomainservices.ClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
