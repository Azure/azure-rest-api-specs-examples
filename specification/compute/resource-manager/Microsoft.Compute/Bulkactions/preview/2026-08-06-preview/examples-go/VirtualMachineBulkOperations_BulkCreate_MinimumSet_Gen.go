package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/VirtualMachineBulkOperations_BulkCreate_MinimumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_BulkCreateOperation_virtualMachineBulkOperationsBulkCreateExampleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkCreateOperation(ctx, "rgBulkactions", "useast2euap", armbulkactions.ExecuteCreateContent{
		ResourceConfigParameters: &armbulkactions.ResourceProvisionPayload{
			ResourceCount: to.Ptr[int32](23),
		},
		ExecutionParameters: &armbulkactions.ExecutionParameters{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkCreateOperationResponse{
	// 	CreateResourceOperationResponse: armbulkactions.CreateResourceOperationResponse{
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("useast2euap"),
	// 		Description: to.Ptr("Bulk create operation"),
	// 	},
	// }
}
