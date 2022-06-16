package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/SecuritySettingsUpdatePost.json
func ExampleDevicesClient_BeginCreateOrUpdateSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdateSecuritySettings(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.SecuritySettings{
			Properties: &armdataboxedge.SecuritySettingsProperties{
				DeviceAdminPassword: &armdataboxedge.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithm("AES256").ToPtr(),
					EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
					Value:                    to.StringPtr("<value>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
