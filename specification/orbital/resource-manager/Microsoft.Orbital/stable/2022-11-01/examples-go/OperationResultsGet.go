package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/OperationResultsGet.json
func ExampleOperationsResultsClient_BeginGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOperationsResultsClient().BeginGet(ctx, "eastus2", "30972f1b-b61d-4fd8-bd34-3dcfa24670f3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResult = armorbital.OperationResult{
	// 	Name: to.Ptr("30972f1b-b61d-4fd8-bd34-3dcfa24670f3"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-10T13:54:55.702Z"); return t}()),
	// 	Error: &armorbital.OperationResultErrorProperties{
	// 		Code: to.Ptr("ReadOnlyPropertyError"),
	// 		Message: to.Ptr("Cannot write to property, AuthorizationStatus, as it is read-only."),
	// 	},
	// 	ID: to.Ptr("https://management.azure.com/Microsoft.Orbital/operationResults/4e2ffff7-b331-4fcb-ab11-b5fa49368188?api-version=2022-11-01"),
	// 	PercentComplete: to.Ptr[float64](1),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-10T13:54:53.655Z"); return t}()),
	// 	Status: to.Ptr(armorbital.StatusFailed),
	// }
}
