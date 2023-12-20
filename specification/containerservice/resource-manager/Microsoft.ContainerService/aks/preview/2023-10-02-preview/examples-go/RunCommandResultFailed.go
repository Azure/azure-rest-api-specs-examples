package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-10-02-preview/examples/RunCommandResultFailed.json
func ExampleManagedClustersClient_GetCommandResult_commandFailedResult() {
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
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Reason: to.Ptr("ImagePullBackoff"),
	// 	},
	// }
}
