package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2022-10-01-preview/examples/ContainerExec.json
func ExampleContainersClient_ExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainersClient().ExecuteCommand(ctx, "demo", "demo1", "container1", armcontainerinstance.ContainerExecRequest{
		Command: to.Ptr("/bin/bash"),
		TerminalSize: &armcontainerinstance.ContainerExecRequestTerminalSize{
			Cols: to.Ptr[int32](12),
			Rows: to.Ptr[int32](12),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ContainerExecResponse = armcontainerinstance.ContainerExecResponse{
	// 	Password: to.Ptr("password"),
	// 	WebSocketURI: to.Ptr("wss://web-socket-uri"),
	// }
}
