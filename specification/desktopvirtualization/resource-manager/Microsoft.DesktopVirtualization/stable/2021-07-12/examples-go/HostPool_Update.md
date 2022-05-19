Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv1.0.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2021-07-12/examples/HostPool_Update.json
func ExampleHostPoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdesktopvirtualization.NewHostPoolsClient("daefabc0-95b4-48b3-b645-8a753a63c4fa", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"resourceGroup1",
		"hostPool1",
		&armdesktopvirtualization.HostPoolsClientUpdateOptions{HostPool: &armdesktopvirtualization.HostPoolPatch{
			Properties: &armdesktopvirtualization.HostPoolPatchProperties{
				Description:                   to.Ptr("des1"),
				FriendlyName:                  to.Ptr("friendly"),
				LoadBalancerType:              to.Ptr(armdesktopvirtualization.LoadBalancerTypeBreadthFirst),
				MaxSessionLimit:               to.Ptr[int32](999999),
				PersonalDesktopAssignmentType: to.Ptr(armdesktopvirtualization.PersonalDesktopAssignmentTypeAutomatic),
				RegistrationInfo: &armdesktopvirtualization.RegistrationInfoPatch{
					ExpirationTime:             to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-01T15:01:54.9571247Z"); return t }()),
					RegistrationTokenOperation: to.Ptr(armdesktopvirtualization.RegistrationTokenOperationUpdate),
				},
				SsoClientID:                 to.Ptr("client"),
				SsoClientSecretKeyVaultPath: to.Ptr("https://keyvault/secret"),
				SsoSecretType:               to.Ptr(armdesktopvirtualization.SSOSecretTypeSharedKey),
				SsoadfsAuthority:            to.Ptr("https://adfs"),
				StartVMOnConnect:            to.Ptr(false),
				VMTemplate:                  to.Ptr("{json:json}"),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
