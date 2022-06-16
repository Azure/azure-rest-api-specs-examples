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
			Kind: armdataboxedge.RoleTypes("IOT").ToPtr(),
			Properties: &armdataboxedge.IoTRoleProperties{
				HostPlatform: armdataboxedge.PlatformType("Linux").ToPtr(),
				IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
					Authentication: &armdataboxedge.Authentication{
						SymmetricKey: &armdataboxedge.SymmetricKey{
							ConnectionString: &armdataboxedge.AsymmetricEncryptedSecret{
								EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithm("AES256").ToPtr(),
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
								EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithm("AES256").ToPtr(),
								EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
								Value:                    to.StringPtr("<value>"),
							},
						},
					},
					DeviceID:   to.StringPtr("<device-id>"),
					IoTHostHub: to.StringPtr("<io-thost-hub>"),
				},
				RoleStatus:    armdataboxedge.RoleStatus("Enabled").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.RolesClientCreateOrUpdateResult)
}
