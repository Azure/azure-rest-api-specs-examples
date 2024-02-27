package armlargeinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_Restart.json
func ExampleAzureLargeInstanceClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureLargeInstanceClient().BeginRestart(ctx, "myResourceGroup", "myALInstance", &armlargeinstance.AzureLargeInstanceClientBeginRestartOptions{ForceParameter: nil})
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
	// res.OperationStatusResult = armlargeinstance.OperationStatusResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	Error: &armlargeinstance.ErrorDetail{
	// 		Code: to.Ptr(""),
	// 		Message: to.Ptr(""),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-04T21:17:24.905Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// }
}
