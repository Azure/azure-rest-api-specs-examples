package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cecf11996f2bfc7958f6d0de814badab23f9720/specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/MarkDevicesShipped.json
func ExampleJobsClient_MarkDevicesShipped() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJobsClient().MarkDevicesShipped(ctx, "TestJobName1", "YourResourceGroupName", armdatabox.MarkDevicesShippedRequest{
		DeliverToDcPackageDetails: &armdatabox.PackageCarrierInfo{
			CarrierName: to.Ptr("testCarrier"),
			TrackingID:  to.Ptr("000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
