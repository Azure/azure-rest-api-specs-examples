package armcontainerservice_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/RunCommandRequest.json
func ExampleManagedClustersClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRunCommand(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.RunCommandRequest{
			ClusterToken: to.StringPtr("<cluster-token>"),
			Command:      to.StringPtr("<command>"),
			Context:      to.StringPtr("<context>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RunCommandResult.ID: %s\n", *res.ID)
}
