package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2021-11-01-preview/examples/ContainerServiceGetOSOptions.json
func ExampleManagedClustersClient_GetOSOptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	res, err := client.GetOSOptions(ctx,
		"<location>",
		&armcontainerservice.ManagedClustersClientGetOSOptionsOptions{ResourceType: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ManagedClustersClientGetOSOptionsResult)
}
