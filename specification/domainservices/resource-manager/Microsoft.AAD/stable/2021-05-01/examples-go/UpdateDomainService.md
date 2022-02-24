Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdomainservices%2Farmdomainservices%2Fv0.2.1/sdk/resourcemanager/domainservices/armdomainservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				ConfigDiagnostics: &armdomainservices.ConfigDiagnostics{
					LastExecuted: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2021-05-05T12:00:23Z;"); return t }()),
					ValidatorResults: []*armdomainservices.ConfigDiagnosticsValidatorResult{
						{
							Issues: []*armdomainservices.ConfigDiagnosticsValidatorResultIssue{
								{
									DescriptionParams: []*string{},
									ID:                to.StringPtr("<id>"),
								}},
							ReplicaSetSubnetDisplayName: to.StringPtr("<replica-set-subnet-display-name>"),
							Status:                      armdomainservices.Status("Warning").ToPtr(),
							ValidatorID:                 to.StringPtr("<validator-id>"),
						}},
				},
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            armdomainservices.NtlmV1("Enabled").ToPtr(),
					SyncNtlmPasswords: armdomainservices.SyncNtlmPasswords("Enabled").ToPtr(),
					TLSV1:             armdomainservices.TLSV1("Disabled").ToPtr(),
				},
				FilteredSync: armdomainservices.FilteredSync("Enabled").ToPtr(),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         armdomainservices.ExternalAccess("Enabled").ToPtr(),
					Ldaps:                  armdomainservices.Ldaps("Enabled").ToPtr(),
					PfxCertificate:         to.StringPtr("<pfx-certificate>"),
					PfxCertificatePassword: to.StringPtr("<pfx-certificate-password>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.StringPtr("jicha@microsoft.com"),
						to.StringPtr("caalmont@microsoft.com")},
					NotifyDcAdmins:     armdomainservices.NotifyDcAdmins("Enabled").ToPtr(),
					NotifyGlobalAdmins: armdomainservices.NotifyGlobalAdmins("Enabled").ToPtr(),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.StringPtr("<location>"),
						SubnetID: to.StringPtr("<subnet-id>"),
					},
					{
						Location: to.StringPtr("<location>"),
						SubnetID: to.StringPtr("<subnet-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientUpdateResult)
}
```
