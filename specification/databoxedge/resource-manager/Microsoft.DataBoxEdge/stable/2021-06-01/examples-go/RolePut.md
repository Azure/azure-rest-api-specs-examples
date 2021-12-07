Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.1.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/RolePut.json
func ExampleRolesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewRolesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<name>",
		"<resource-group-name>",
		&armdataboxedge.IoTRole{
			Role: armdataboxedge.Role{
				Kind: armdataboxedge.RoleTypesIOT.ToPtr(),
			},
			Properties: &armdataboxedge.IoTRoleProperties{
				HostPlatform: armdataboxedge.PlatformTypeLinux.ToPtr(),
				IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
					Authentication: &armdataboxedge.Authentication{
						SymmetricKey: &armdataboxedge.SymmetricKey{
							ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
								EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithmAES256.ToPtr(),
								EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
								Value:                    to.StringPtr("<value>"),
							},
						},
					},
					DeviceID:   to.StringPtr("<device-id>"),
					IoTHostHub: to.StringPtr("<io-thost-hub>"),
				},
				IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
					Authentication: &armdataboxedge.Authentication{
						SymmetricKey: &armdataboxedge.SymmetricKey{
							ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
								EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithmAES256.ToPtr(),
								EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
								Value:                    to.StringPtr("<value>"),
							},
						},
					},
					DeviceID:   to.StringPtr("<device-id>"),
					IoTHostHub: to.StringPtr("<io-thost-hub>"),
				},
				RoleStatus:    armdataboxedge.RoleStatusEnabled.ToPtr(),
				ShareMappings: []*armdataboxedge.MountPointMap{},
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
	log.Printf("RoleClassification.GetRole().ID: %s\n", *res.GetRole().ID)
}
```
