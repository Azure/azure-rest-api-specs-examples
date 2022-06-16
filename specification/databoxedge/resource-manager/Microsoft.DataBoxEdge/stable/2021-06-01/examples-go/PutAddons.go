package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/PutAddons.json
func ExampleAddonsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewAddonsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<role-name>",
		"<addon-name>",
		"<resource-group-name>",
		&armdataboxedge.ArcAddon{
			Kind: armdataboxedge.AddonType("ArcForKubernetes").ToPtr(),
			Properties: &armdataboxedge.ArcAddonProperties{
				ResourceGroupName: to.StringPtr("<resource-group-name>"),
				ResourceLocation:  to.StringPtr("<resource-location>"),
				ResourceName:      to.StringPtr("<resource-name>"),
				SubscriptionID:    to.StringPtr("<subscription-id>"),
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
	log.Printf("Response result: %#v\n", res.AddonsClientCreateOrUpdateResult)
}
