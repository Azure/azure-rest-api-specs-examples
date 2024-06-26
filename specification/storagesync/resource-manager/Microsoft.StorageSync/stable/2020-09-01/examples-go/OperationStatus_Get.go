package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/OperationStatus_Get.json
func ExampleOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusClient().Get(ctx, "SampleResourceGroup_1", "westus", "828219ea-083e-48b5-89ea-8fd9991b2e75", "14b50e24-f68d-4b29-a882-38be9dfb8bd1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armstoragesync.OperationStatus{
	// 	Name: to.Ptr("14b50e24-f68d-4b29-a882-38be9dfb8bd1"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T23:53:23.373Z"); return t}()),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T23:53:21.732Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
