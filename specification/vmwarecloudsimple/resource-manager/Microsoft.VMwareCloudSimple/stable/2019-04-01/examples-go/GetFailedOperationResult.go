package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetFailedOperationResult.json
func ExampleOperationsClient_Get_getFailedOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().Get(ctx, "westus2", "https://management.azure.com/", "d030bb3f-7d53-11e9-8e09-9a86872085ff", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResource = armvmwarecloudsimple.OperationResource{
	// 	Name: to.Ptr("d030bb3f-7d53-11e9-8e09-9a86872085ff"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-23T12:11:00.655Z"); return t}()),
	// 	Error: &armvmwarecloudsimple.OperationError{
	// 		Code: to.Ptr("InternalError"),
	// 		Message: to.Ptr("Internal Service error"),
	// 	},
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/microsoft.vmwarecloudsimple/locations/westus2/operationresults/d030bb3f-7d53-11e9-8e09-9a86872085ff"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-23T12:10:54.012Z"); return t}()),
	// 	Status: to.Ptr("Failed"),
	// }
}
