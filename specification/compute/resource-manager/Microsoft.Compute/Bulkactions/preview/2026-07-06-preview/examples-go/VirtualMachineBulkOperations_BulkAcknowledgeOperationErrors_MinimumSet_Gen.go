package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/VirtualMachineBulkOperations_BulkAcknowledgeOperationErrors_MinimumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_BulkAcknowledgeOperationErrors_virtualMachineBulkOperationsBulkAcknowledgeOperationErrorsMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkAcknowledgeOperationErrors(ctx, "rgBulkactions", "useast2euap", armbulkactions.AcknowledgeBulkOperationErrorsRequest{
		OperationIDs: []*string{
			to.Ptr("af449548-8e1a-4079-874e-2caa4ff783cc"),
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
	// 			to.Ptr("af449548-8e1a-4079-874e-2caa4ff783cc"),
	// 		},
	// 		NotFound: []*string{
	// 			to.Ptr("bf559658-9f2b-5180-985f-3dbb5678901d"),
	// 		},
	// 		Skipped: []*string{
	// 			to.Ptr("cf669768-af3c-6291-a96g-4ecc6789012e"),
	// 		},
	// 	},
	// }
}
