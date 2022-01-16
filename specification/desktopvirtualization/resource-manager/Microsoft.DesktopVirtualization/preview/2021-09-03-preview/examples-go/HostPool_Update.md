Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.2.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/HostPool_Update.json
func ExampleHostPoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewHostPoolsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<host-pool-name>",
		&armdesktopvirtualization.HostPoolsClientUpdateOptions{HostPool: &armdesktopvirtualization.HostPoolPatch{
			Properties: &armdesktopvirtualization.HostPoolPatchProperties{
				Description:                   to.StringPtr("<description>"),
				FriendlyName:                  to.StringPtr("<friendly-name>"),
				LoadBalancerType:              armdesktopvirtualization.LoadBalancerType("BreadthFirst").ToPtr(),
				MaxSessionLimit:               to.Int32Ptr(999999),
				PersonalDesktopAssignmentType: armdesktopvirtualization.PersonalDesktopAssignmentType("Automatic").ToPtr(),
				PublicNetworkAccess:           armdesktopvirtualization.PublicNetworkAccess("Enabled").ToPtr(),
				RegistrationInfo: &armdesktopvirtualization.RegistrationInfoPatch{
					ExpirationTime:             to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-01T15:01:54.9571247Z"); return t }()),
					RegistrationTokenOperation: armdesktopvirtualization.RegistrationTokenOperation("Update").ToPtr(),
				},
				SsoClientID:                 to.StringPtr("<sso-client-id>"),
				SsoClientSecretKeyVaultPath: to.StringPtr("<sso-client-secret-key-vault-path>"),
				SsoSecretType:               armdesktopvirtualization.SSOSecretType("SharedKey").ToPtr(),
				SsoadfsAuthority:            to.StringPtr("<ssoadfs-authority>"),
				StartVMOnConnect:            to.BoolPtr(false),
				VMTemplate:                  to.StringPtr("<vmtemplate>"),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.HostPoolsClientUpdateResult)
}
```
