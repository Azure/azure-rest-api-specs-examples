package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2024-10-01-preview/examples/OperationProgress_Get_BackupAndExport.json
func ExampleOperationProgressClient_Get_operationProgressGetBackupAndExport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationProgressClient().Get(ctx, "westus", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationProgressResult = armmysqlflexibleservers.OperationProgressResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.DBforMySQL/locations/westus/operationProgress/00000000-0000-0000-0000-000000000000"),
	// 	PercentComplete: to.Ptr[float32](10),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-29T07:34:02.328Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// 	Properties: &armmysqlflexibleservers.BackupAndExportResponseType{
	// 		ObjectType: to.Ptr(armmysqlflexibleservers.ObjectTypeBackupAndExportResponse),
	// 		BackupMetadata: to.Ptr("{\"key1\":\"value1\",\"key2\":\"value2\"}"),
	// 		DataTransferredInBytes: to.Ptr[int64](102),
	// 		DatasourceSizeInBytes: to.Ptr[int64](1024),
	// 	},
	// }
}
