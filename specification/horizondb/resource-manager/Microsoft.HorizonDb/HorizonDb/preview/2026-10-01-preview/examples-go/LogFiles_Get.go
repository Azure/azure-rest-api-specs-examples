package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/LogFiles_Get.json
func ExampleLogFilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogFilesClient().Get(ctx, "example-resource-group", "example-cluster", "11111111-1111-1111-1111-111111111111", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.LogFilesClientGetResponse{
	// 	LogFile: armhorizondb.LogFile{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/logFiles/11111111-1111-1111-1111-111111111111"),
	// 		Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/logFiles"),
	// 		Properties: &armhorizondb.LogFileProperties{
	// 			LogType: to.Ptr(armhorizondb.LogFileTypePostgreSQLServer),
	// 			CreatedAtUTC: to.Ptr(time.Date(2026, time.September, 11, 12, 0, 0, 0, time.UTC)),
	// 			LastModifiedAtUTC: to.Ptr(time.Date(2026, time.September, 11, 12, 15, 0, 0, time.UTC)),
	// 			RetentionExpiresAtUTC: to.Ptr(time.Date(2026, time.September, 25, 12, 0, 0, 0, time.UTC)),
	// 			IsMutable: to.Ptr(true),
	// 			SizeInKilobytes: to.Ptr[int64](2048),
	// 			Pool: to.Ptr("example-pool-01"),
	// 			Node: to.Ptr("node-01"),
	// 			Role: to.Ptr(armhorizondb.NodeRoleReadWrite),
	// 		},
	// 	},
	// }
}
