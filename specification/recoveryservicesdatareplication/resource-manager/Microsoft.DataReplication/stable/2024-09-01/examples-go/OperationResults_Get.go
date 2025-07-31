package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/OperationResults_Get.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationResultsClient().Get(ctx, "rgswagger_2024-09-01", "lghle", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.OperationResultsClientGetResponse{
	// 	OperationStatus: &armrecoveryservicesdatareplication.OperationStatus{
	// 		ID: to.Ptr("inhnjrxtqtyumflbzdne"),
	// 		Name: to.Ptr("eawsgzqm"),
	// 		Status: to.Ptr("jjijbsrjfqvqzriekxfvynb"),
	// 		StartTime: to.Ptr("slmhzgrfzkkrxdalacmidyxijq"),
	// 		EndTime: to.Ptr("ixkalnbiajfpjtur"),
	// 	},
	// }
}
