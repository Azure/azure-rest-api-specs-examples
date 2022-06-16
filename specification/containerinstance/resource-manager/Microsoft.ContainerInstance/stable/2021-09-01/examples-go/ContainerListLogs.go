package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"
)

// x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerListLogs.json
func ExampleContainersClient_ListLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerinstance.NewContainersClient("<subscription-id>", cred, nil)
	_, err = client.ListLogs(ctx,
		"<resource-group-name>",
		"<container-group-name>",
		"<container-name>",
		&armcontainerinstance.ContainersListLogsOptions{Tail: to.Int32Ptr(10),
			Timestamps: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
}
