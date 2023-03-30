package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreateGeoRestoreMode.json
func ExampleServersClient_BeginCreate_createAServerAsAGeoRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServersClient().BeginCreate(ctx, "TargetResourceGroup", "targetserver", armmysql.ServerForCreate{
		Location: to.Ptr("westus"),
		Properties: &armmysql.ServerPropertiesForGeoRestore{
			CreateMode:     to.Ptr(armmysql.CreateModeGeoRestore),
			SourceServerID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/servers/sourceserver"),
		},
		SKU: &armmysql.SKU{
			Name:     to.Ptr("GP_Gen5_2"),
			Capacity: to.Ptr[int32](2),
			Family:   to.Ptr("Gen5"),
			Tier:     to.Ptr(armmysql.SKUTierGeneralPurpose),
		},
		Tags: map[string]*string{
			"ElasticServer": to.Ptr("1"),
		},
	}, nil)
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
	// res.Server = armmysql.Server{
	// 	Name: to.Ptr("targetserver"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/targetserver"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"elasticServer": to.Ptr("1"),
	// 	},
	// 	Properties: &armmysql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("cloudsa"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-14T18:02:41.577+00:00"); return t}()),
	// 		FullyQualifiedDomainName: to.Ptr("targetserver.mysql.database.azure.com"),
	// 		SSLEnforcement: to.Ptr(armmysql.SSLEnforcementEnumEnabled),
	// 		StorageProfile: &armmysql.StorageProfile{
	// 			BackupRetentionDays: to.Ptr[int32](14),
	// 			GeoRedundantBackup: to.Ptr(armmysql.GeoRedundantBackupEnabled),
	// 			StorageMB: to.Ptr[int32](128000),
	// 		},
	// 		UserVisibleState: to.Ptr(armmysql.ServerStateReady),
	// 		Version: to.Ptr(armmysql.ServerVersionFive7),
	// 	},
	// 	SKU: &armmysql.SKU{
	// 		Name: to.Ptr("GP_Gen5_2"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Family: to.Ptr("Gen5"),
	// 		Tier: to.Ptr(armmysql.SKUTierGeneralPurpose),
	// 	},
	// }
}
