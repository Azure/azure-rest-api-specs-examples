package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/OperationStatus_Get.json
func ExampleOperationStatusesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusesClient().Get(ctx, "westus3", "3fa1a29d-e807-488d-81d1-f1c5456a08cd", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armdevcenter.OperationStatus{
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T15:16:55.100Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/providers/Microsoft.DevCenter/locations/CENTRALUS/operationStatuses/722e7bc4-60fa-4e6b-864f-d5bf12b9adc4"),
	// 	PercentComplete: to.Ptr[float32](100),
	// 	ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T15:16:29.500Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
