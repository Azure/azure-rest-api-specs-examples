package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetOperationResult.json
func ExampleOperationsClient_Get_getOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().Get(ctx, "westus2", "https://management.azure.com/", "f8e1c8f1-7d52-11e9-8e07-9a86872085ff", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationResource = armvmwarecloudsimple.OperationResource{
	// 	Name: to.Ptr("f8e1c8f1-7d52-11e9-8e07-9a86872085ff"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-23T12:05:55.660Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/microsoft.vmwarecloudsimple/locations/westus2/operationresults/f8e1c8f1-7d52-11e9-8e07-9a86872085ff"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-23T12:04:52.784Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
