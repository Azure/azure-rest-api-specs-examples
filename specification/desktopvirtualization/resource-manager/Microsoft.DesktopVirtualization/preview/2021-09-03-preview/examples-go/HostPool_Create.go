package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_Create.json
func ExampleHostPoolsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewHostPoolsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		armdesktopvirtualization.HostPool{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armdesktopvirtualization.HostPoolProperties{
				Description:      to.StringPtr("<description>"),
				FriendlyName:     to.StringPtr("<friendly-name>"),
				HostPoolType:     armdesktopvirtualization.HostPoolType("Pooled").ToPtr(),
				LoadBalancerType: armdesktopvirtualization.LoadBalancerType("BreadthFirst").ToPtr(),
				MaxSessionLimit:  to.Int32Ptr(999999),
				MigrationRequest: &armdesktopvirtualization.MigrationRequestProperties{
					MigrationPath: to.StringPtr("<migration-path>"),
					Operation:     armdesktopvirtualization.Operation("Start").ToPtr(),
				},
				PersonalDesktopAssignmentType: armdesktopvirtualization.PersonalDesktopAssignmentType("Automatic").ToPtr(),
				PreferredAppGroupType:         armdesktopvirtualization.PreferredAppGroupType("Desktop").ToPtr(),
				PublicNetworkAccess:           armdesktopvirtualization.PublicNetworkAccess("Enabled").ToPtr(),
				RegistrationInfo: &armdesktopvirtualization.RegistrationInfo{
					ExpirationTime:             to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-01T14:01:54.9571247Z"); return t }()),
					RegistrationTokenOperation: armdesktopvirtualization.RegistrationTokenOperation("Update").ToPtr(),
				},
				SsoClientID:                 to.StringPtr("<sso-client-id>"),
				SsoClientSecretKeyVaultPath: to.StringPtr("<sso-client-secret-key-vault-path>"),
				SsoSecretType:               armdesktopvirtualization.SSOSecretType("SharedKey").ToPtr(),
				SsoadfsAuthority:            to.StringPtr("<ssoadfs-authority>"),
				StartVMOnConnect:            to.BoolPtr(false),
				VMTemplate:                  to.StringPtr("<vmtemplate>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HostPoolsClientCreateOrUpdateResult)
}
