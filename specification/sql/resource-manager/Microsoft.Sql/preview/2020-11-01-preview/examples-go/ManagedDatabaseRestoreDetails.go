package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseRestoreDetails.json
func ExampleManagedDatabaseRestoreDetailsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseRestoreDetailsClient().Get(ctx, "Default-SQL-SouthEastAsia", "managedInstance", "testdb", armsql.RestoreDetailsNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedDatabaseRestoreDetailsResult = armsql.ManagedDatabaseRestoreDetailsResult{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/restoreDetails"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/managedInstance/databases/testdb/restoreDetails/current"),
	// 	Properties: &armsql.ManagedDatabaseRestoreDetailsProperties{
	// 		CurrentRestoringFileName: to.Ptr("full_0"),
	// 		LastUploadedFileName: to.Ptr("log2_3"),
	// 		LastUploadedFileTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T11:38:33Z"); return t}()),
	// 		NumberOfFilesDetected: to.Ptr[int64](26),
	// 		PercentCompleted: to.Ptr[float64](0),
	// 		Status: to.Ptr("Restoring"),
	// 		UnrestorableFiles: []*string{
	// 		},
	// 	},
	// }
}
