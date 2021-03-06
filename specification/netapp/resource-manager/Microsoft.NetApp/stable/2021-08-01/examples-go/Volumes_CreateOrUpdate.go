package armnetapp_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Volumes_CreateOrUpdate.json
func ExampleVolumesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewVolumesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pool-name>",
		"<volume-name>",
		armnetapp.Volume{
			Location: to.StringPtr("<location>"),
			Properties: &armnetapp.VolumeProperties{
				CreationToken:       to.StringPtr("<creation-token>"),
				EncryptionKeySource: to.StringPtr("<encryption-key-source>"),
				ServiceLevel:        armnetapp.ServiceLevel("Premium").ToPtr(),
				SubnetID:            to.StringPtr("<subnet-id>"),
				ThroughputMibps:     to.Float32Ptr(128),
				UsageThreshold:      to.Int64Ptr(107374182400),
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
	log.Printf("Response result: %#v\n", res.VolumesClientCreateOrUpdateResult)
}
