package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2021-12-01/examples/MarkDevicesShipped.json
func ExampleJobsClient_MarkDevicesShipped() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.MarkDevicesShipped(ctx,
		"<job-name>",
		"<resource-group-name>",
		armdatabox.MarkDevicesShippedRequest{
			DeliverToDcPackageDetails: &armdatabox.PackageCarrierInfo{
				CarrierName: to.Ptr("<carrier-name>"),
				TrackingID:  to.Ptr("<tracking-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
