package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/RegionConfigurationByResourceGroup.json
func ExampleServiceClient_RegionConfigurationByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	res, err := client.RegionConfigurationByResourceGroup(ctx,
		"<resource-group-name>",
		"<location>",
		armdatabox.RegionConfigurationRequest{
			ScheduleAvailabilityRequest: &armdatabox.ScheduleAvailabilityRequest{
				SKUName:         armdatabox.SKUNameDataBox.ToPtr(),
				StorageLocation: to.StringPtr("<storage-location>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServiceClientRegionConfigurationByResourceGroupResult)
}
