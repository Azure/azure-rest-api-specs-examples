package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveoutput-create.json
func ExampleLiveOutputsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveOutputsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		"<live-output-name>",
		armmediaservices.LiveOutput{
			Properties: &armmediaservices.LiveOutputProperties{
				Description:         to.StringPtr("<description>"),
				ArchiveWindowLength: to.StringPtr("<archive-window-length>"),
				AssetName:           to.StringPtr("<asset-name>"),
				Hls: &armmediaservices.Hls{
					FragmentsPerTsSegment: to.Int32Ptr(5),
				},
				ManifestName: to.StringPtr("<manifest-name>"),
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
	log.Printf("Response result: %#v\n", res.LiveOutputsClientCreateResult)
}
