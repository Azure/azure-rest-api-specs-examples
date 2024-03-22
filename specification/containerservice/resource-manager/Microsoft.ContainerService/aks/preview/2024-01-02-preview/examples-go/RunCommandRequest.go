package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44a81f85be1e4797fbf5e290fc6b41d48788a6ba/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-01-02-preview/examples/RunCommandRequest.json
func ExampleManagedClustersClient_BeginRunCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginRunCommand(ctx, "rg1", "clustername1", armcontainerservice.RunCommandRequest{
		ClusterToken: to.Ptr(""),
		Command:      to.Ptr("kubectl apply -f ns.yaml"),
		Context:      to.Ptr(""),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RunCommandResult = armcontainerservice.RunCommandResult{
	// 	ID: to.Ptr("def7b3ea71bd4f7e9d226ddbc0f00ad9"),
	// 	Properties: &armcontainerservice.CommandResultProperties{
	// 		ExitCode: to.Ptr[int32](0),
	// 		FinishedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-17T00:28:33.000Z"); return t}()),
	// 		Logs: to.Ptr("namespace dummy created"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-17T00:28:20.000Z"); return t}()),
	// 	},
	// }
}
