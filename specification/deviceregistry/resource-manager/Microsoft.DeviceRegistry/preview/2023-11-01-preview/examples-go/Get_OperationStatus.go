package armdeviceregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c77bbf822be2deaac1b690270c6cd03a52df0e37/specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Get_OperationStatus.json
func ExampleOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusClient().Get(ctx, "testLocation", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armdeviceregistry.OperationStatusResult{
	// 	Name: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-21T13:28:03.898Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DeviceRegistry/locations/testLocation/operationStatuses/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	PercentComplete: to.Ptr[float32](100),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-21T13:27:03.898Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
