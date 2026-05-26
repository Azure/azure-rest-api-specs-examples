package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v3"
)

// Generated from example definition: 2025-07-01/MarkDevicesShipped.json
func ExampleJobsClient_MarkDevicesShipped() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("YourSubscriptionId", cred, nil)
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
