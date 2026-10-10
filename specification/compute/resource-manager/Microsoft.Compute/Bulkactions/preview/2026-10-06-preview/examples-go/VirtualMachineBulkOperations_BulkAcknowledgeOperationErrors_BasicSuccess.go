package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkAcknowledgeOperationErrors_BasicSuccess.json
func ExampleVirtualMachineBulkOperationsClient_BulkAcknowledgeOperationErrors_oneAcknowledgeMultipleOperationErrors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkAcknowledgeOperationErrors(ctx, "example-rg", "eastus", armbulkactions.AcknowledgeBulkOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			to.Ptr("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkAcknowledgeOperationErrorsResponse{
	// 	AcknowledgeBulkOperationErrorsResponse: armbulkactions.AcknowledgeBulkOperationErrorsResponse{
	// 		Acknowledged: []*string{
	// 			to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
	// 			to.Ptr("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
	// 		},
	// 		NotFound: []*string{
	// 		},
	// 		Skipped: []*string{
	// 		},
	// 	},
	// }
}
