package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_SqlDBTable_Create.json
func ExampleDataSetsClient_Create_dataSetsSqlDbTableCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetsClient().Create(ctx, "SampleResourceGroup", "Account1", "Share1", "Dataset1", &armdatashare.SQLDBTableDataSet{
		Kind: to.Ptr(armdatashare.DataSetKindSQLDBTable),
		Properties: &armdatashare.SQLDBTableProperties{
			DatabaseName:        to.Ptr("SqlDB1"),
			SchemaName:          to.Ptr("dbo"),
			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
			TableName:           to.Ptr("Table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetsClientCreateResponse{
	// 	                            DataSetClassification: &armdatashare.SQLDBTableDataSet{
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shares/dataSets"),
	// 		ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1/datasets/Dataset1"),
	// 		Kind: to.Ptr(armdatashare.DataSetKindSQLDBTable),
	// 		Properties: &armdatashare.SQLDBTableProperties{
	// 			DataSetID: to.Ptr("a08f184b-0567-4b11-ba22-a1199336d226"),
	// 			DatabaseName: to.Ptr("SqlDB1"),
	// 			SchemaName: to.Ptr("dbo"),
	// 			SQLServerResourceID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"),
	// 			TableName: to.Ptr("Table1"),
	// 		},
	// 	},
	// 	                        }
}
