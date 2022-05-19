Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv1.0.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RolePut.json
func ExampleRolesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewRolesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testedgedevice",
		"IoTRole1",
		"GroupForEdgeAutomation",
		&armdataboxedge.IoTRole{
			Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
			Properties: &armdataboxedge.IoTRoleProperties{
				HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
				IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
					Authentication: &armdataboxedge.Authentication{
						SymmetricKey: &armdataboxedge.SymmetricKey{
							ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
								EncryptionAlgorithm:      to.Ptr(armdataboxedge.EncryptionAlgorithmAES256),
								EncryptionCertThumbprint: to.Ptr("348586569999244"),
								Value:                    to.Ptr("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotDevice;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>"),
							},
						},
					},
					DeviceID:   to.Ptr("iotdevice"),
					IoTHostHub: to.Ptr("iothub.azure-devices.net"),
				},
				IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
					Authentication: &armdataboxedge.Authentication{
						SymmetricKey: &armdataboxedge.SymmetricKey{
							ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
								EncryptionAlgorithm:      to.Ptr(armdataboxedge.EncryptionAlgorithmAES256),
								EncryptionCertThumbprint: to.Ptr("1245475856069999244"),
								Value:                    to.Ptr("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotEdge;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>"),
							},
						},
					},
					DeviceID:   to.Ptr("iotEdge"),
					IoTHostHub: to.Ptr("iothub.azure-devices.net"),
				},
				RoleStatus:    to.Ptr(armdataboxedge.RoleStatusEnabled),
				ShareMappings: []*armdataboxedge.MountPointMap{},
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
```
