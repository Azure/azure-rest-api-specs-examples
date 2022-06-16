package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/TriggerPut.json
func ExampleTriggersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewTriggersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<name>",
		"<resource-group-name>",
		&armdataboxedge.FileEventTrigger{
			Kind: armdataboxedge.TriggerEventType("FileEvent").ToPtr(),
			Properties: &armdataboxedge.FileTriggerProperties{
				CustomContextTag: to.StringPtr("<custom-context-tag>"),
				SinkInfo: &armdataboxedge.RoleSinkInfo{
					RoleID: to.StringPtr("<role-id>"),
				},
				SourceInfo: &armdataboxedge.FileSourceInfo{
					ShareID: to.StringPtr("<share-id>"),
				},
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
	log.Printf("Response result: %#v\n", res.TriggersClientCreateOrUpdateResult)
}
