package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2017-07-01/examples/ChildAvailabilityStatus_GetByResource.json
func ExampleChildAvailabilityStatusesClient_GetByResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armresourcehealth.NewChildAvailabilityStatusesClient(cred, nil)
	res, err := client.GetByResource(ctx,
		"<resource-uri>",
		&armresourcehealth.ChildAvailabilityStatusesClientGetByResourceOptions{Filter: nil,
			Expand: to.StringPtr("<expand>"),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ChildAvailabilityStatusesClientGetByResourceResult)
}
