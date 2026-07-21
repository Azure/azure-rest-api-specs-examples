package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/VirtualEnclave_HandleApprovalDeletion.json
func ExampleVirtualEnclaveClient_BeginHandleApprovalDeletion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualEnclaveClient().BeginHandleApprovalDeletion(ctx, "rgopenapi", "TestMyEnclave", armenclave.ApprovalDeletionCallbackRequest{
		ResourceRequestAction: to.Ptr(armenclave.ApprovalDeletionCallbackRequestResourceRequestActionCreate),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.VirtualEnclaveClientHandleApprovalDeletionResponse{
	// 	ApprovalActionResponse: armenclave.ApprovalActionResponse{
	// 		Message: to.Ptr("Approval state change handled successfully."),
	// 	},
	// }
}
