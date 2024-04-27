package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1dd99306d14fd6c602f47652a209a4a6812c368c/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-02-02-preview/examples/RunCommandResultSucceed.json
func ExampleManagedClustersClient_GetCommandResult_commandSucceedResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().GetCommandResult(ctx, "rg1", "clustername1", "def7b3ea71bd4f7e9d226ddbc0f00ad9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
