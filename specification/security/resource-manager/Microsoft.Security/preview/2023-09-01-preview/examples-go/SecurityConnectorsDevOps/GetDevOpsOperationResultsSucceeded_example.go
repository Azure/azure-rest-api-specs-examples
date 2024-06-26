package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetDevOpsOperationResultsSucceeded_example.json
func ExampleDevOpsOperationResultsClient_Get_getDevOpsOperationResultsSucceeded() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevOpsOperationResultsClient().Get(ctx, "myRg", "mySecurityConnectorName", "4e826cf1-5c36-4808-a7d2-fb4f5170978b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armsecurity.OperationStatusResult{
	// 	Name: to.Ptr("4e826cf1-5c36-4808-a7d2-fb4f5170978b"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-21T03:25:15.000Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/operationResults/4e826cf1-5c36-4808-a7d2-fb4f5170978b?api-version=2023-09-01-preview"),
	// 	PercentComplete: to.Ptr[float32](100),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-21T03:23:15.000Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
