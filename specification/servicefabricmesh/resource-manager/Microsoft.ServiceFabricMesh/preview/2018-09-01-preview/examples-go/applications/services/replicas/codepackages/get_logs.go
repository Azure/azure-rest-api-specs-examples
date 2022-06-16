package armservicefabricmesh_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabricmesh/resource-manager/Microsoft.ServiceFabricMesh/preview/2018-09-01-preview/examples/applications/services/replicas/codepackages/get_logs.json
func ExampleCodePackageClient_GetContainerLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabricmesh.NewCodePackageClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetContainerLogs(ctx,
		"sbz_demo",
		"sbzDocApp",
		"sbzDocService",
		"0",
		"sbzDocCode",
		&armservicefabricmesh.CodePackageClientGetContainerLogsOptions{Tail: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
