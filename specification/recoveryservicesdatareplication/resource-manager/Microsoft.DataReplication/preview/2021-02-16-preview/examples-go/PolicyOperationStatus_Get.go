package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/PolicyOperationStatus_Get.json
func ExamplePolicyOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyOperationStatusClient().Get(ctx, "rgrecoveryservicesdatareplication", "4", "xczxcwec", "wdqfsdxv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armrecoveryservicesdatareplication.OperationStatus{
	// 	Name: to.Ptr("wzdasptnwlxgobklayoqapjcgcf"),
	// 	EndTime: to.Ptr("nauyrfh"),
	// 	ID: to.Ptr("sf"),
	// 	StartTime: to.Ptr("xuzwmfrhluafmwwsmzqxsytyehsh"),
	// 	Status: to.Ptr("plbnngzfppdram"),
	// }
}
