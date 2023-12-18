package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_Shutdown.json
func ExampleAzureBareMetalInstancesClient_BeginShutdown() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbaremetalinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureBareMetalInstancesClient().BeginShutdown(ctx, "myResourceGroup", "myABMInstance", nil)
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
	// res.OperationStatus = armbaremetalinfrastructure.OperationStatus{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	Error: &armbaremetalinfrastructure.OperationStatusError{
	// 		Code: to.Ptr(""),
	// 		Message: to.Ptr(""),
	// 	},
	// 	StartTime: to.Ptr("2023-08-01T21:17:24.9052926Z"),
	// 	Status: to.Ptr(armbaremetalinfrastructure.AsyncOperationStatus("InProgress")),
	// }
}
