package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/RefreshExternalGovernanceStatus.json
func ExampleServersClient_BeginRefreshStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginRefreshStatus(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", nil)
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
	// res.RefreshExternalGovernanceStatusOperationResult = armsql.RefreshExternalGovernanceStatusOperationResult{
	// 	Name: to.Ptr("9d9a794a-5cec-4f23-af70-d29511b522a4"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/refreshExternalGovernanceStatusOperationResults"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.Sql/locations/japaneast/refreshExternalGovernanceStatusOperationResults/9d9a794a-5cec-4f23-af70-d29511b522a4"),
	// 	Properties: &armsql.RefreshExternalGovernanceStatusOperationResultProperties{
	// 		QueuedTime: to.Ptr("2/12/2022 8:33:27 PM"),
	// 		RequestID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		RequestType: to.Ptr("UpdatePurviewMetadata"),
	// 		ServerName: to.Ptr("testsvr.database.windows.net"),
	// 		Status: to.Ptr("Completed"),
	// 	},
	// }
}
