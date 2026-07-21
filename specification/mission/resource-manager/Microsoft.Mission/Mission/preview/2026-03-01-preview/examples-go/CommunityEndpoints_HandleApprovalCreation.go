package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/CommunityEndpoints_HandleApprovalCreation.json
func ExampleCommunityEndpointsClient_BeginHandleApprovalCreation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("73CEECEF-2C30-488E-946F-D20F414D99BA", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunityEndpointsClient().BeginHandleApprovalCreation(ctx, "rgopenapi", "TestMyCommunity", "TestMyCommunityEndpoint", armenclave.ApprovalCallbackRequest{
		ResourceRequestAction: to.Ptr(armenclave.ApprovalCallbackRequestResourceRequestActionCreate),
		ApprovalStatus:        to.Ptr(armenclave.ApprovalCallbackRequestApprovalStatusApproved),
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
	// res = armenclave.CommunityEndpointsClientHandleApprovalCreationResponse{
	// 	ApprovalActionResponse: armenclave.ApprovalActionResponse{
	// 		Message: to.Ptr("Approval state change handled successfully."),
	// 	},
	// }
}
