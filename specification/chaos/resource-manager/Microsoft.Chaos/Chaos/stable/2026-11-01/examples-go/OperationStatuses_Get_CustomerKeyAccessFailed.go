package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-11-01/OperationStatuses_Get_CustomerKeyAccessFailed.json
func ExampleOperationStatusesClient_Get_getAFailedWorkspaceEncryptionOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusesClient().Get(ctx, "westus2", "4bdadd97-207c-4de8-9bba-08339ae099c7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.OperationStatusesClientGetResponse{
	// 	OperationStatusResult: armchaos.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/providers/Microsoft.Chaos/locations/westus2/operationStatuses/4bdadd97-207c-4de8-9bba-08339ae099c7"),
	// 		Name: to.Ptr("4bdadd97-207c-4de8-9bba-08339ae099c7"),
	// 		Status: to.Ptr("Failed"),
	// 		StartTime: to.Ptr(time.Date(2026, time.November, 1, 11, 59, 0, 0, time.UTC)),
	// 		EndTime: to.Ptr(time.Date(2026, time.November, 1, 12, 0, 0, 0, time.UTC)),
	// 		Error: &armchaos.ErrorDetail{
	// 			Code: to.Ptr("CustomerKeyAccessFailed"),
	// 			Message: to.Ptr("Azure Storage could not use the customer key. Restore access to the required key material and retry the Workspace update. The encryption mode cannot be changed after Workspace creation."),
	// 			Target: to.Ptr("properties.encryption.customerManagedKeyEncryption.keyEncryptionKeyUrl"),
	// 		},
	// 	},
	// }
}
