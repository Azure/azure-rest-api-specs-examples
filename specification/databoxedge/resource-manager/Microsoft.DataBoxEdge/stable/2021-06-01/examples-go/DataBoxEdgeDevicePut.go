package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePut.json
func ExampleDevicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.Device{
			Location: to.StringPtr("<location>"),
			SKU: &armdataboxedge.SKUInfo{
				Name: armdataboxedge.SKUName("Edge").ToPtr(),
				Tier: armdataboxedge.SKUTier("Standard").ToPtr(),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientCreateOrUpdateResult)
}
