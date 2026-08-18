package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices/v3"
)

// Generated from example definition: 2026-07-01/DeletedVaults_GetOperationStatus.json
func ExampleDeletedVaultsClient_GetOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservices.NewClientFactory("77777777-b0c6-47a2-b37c-d8e65a629c18", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedVaultsClient().GetOperationStatus(ctx, "westus", "swaggerExample", "YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA==", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservices.DeletedVaultsClientGetOperationStatusResponse{
	// 	OperationResource: armrecoveryservices.OperationResource{
	// 		ID: to.Ptr("/subscriptions/77777777-b0c6-47a2-b37c-d8e65a629c18/providers/Microsoft.RecoveryServices/locations/westus/deletedVaults/swaggerExample/operations/YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA=="),
	// 		Name: to.Ptr("YWUzNDFkMzQtZmM5OS00MmUyLWEzNDMtZGJkMDIxZjlmZjgzOzdmYzBiMzhmLTc2NmItNDM5NS05OWQ1LTVmOGEzNzg4MWQzNA=="),
	// 		Status: to.Ptr("Succeeded"),
	// 		StartTime: to.Ptr(time.Date(2024, time.September, 20, 9, 49, 44, 47849600, time.UTC)),
	// 		EndTime: to.Ptr(time.Date(2024, time.September, 20, 9, 49, 46, 0, time.UTC)),
	// 	},
	// }
}
