package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSets_SynapseWorkspaceSqlPoolTable_Create.json
func ExampleDataSetsClient_Create_dataSetsSynapseWorkspaceSqlPoolTableCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetsClient().Create(ctx, "SampleResourceGroup", "sourceAccount", "share1", "dataset1", &armdatashare.SynapseWorkspaceSQLPoolTableDataSet{
		Kind: to.Ptr(armdatashare.DataSetKindSynapseWorkspaceSQLPoolTable),
		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetProperties{
			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetsClientCreateResponse{
	// 	                            DataSetClassification: &armdatashare.SynapseWorkspaceSQLPoolTableDataSet{
	// 		Name: to.Ptr("dataset1"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shares/dataSets"),
	// 		ID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/sourceAccount/shares/share1/dataSets/dataset1"),
	// 		Kind: to.Ptr(armdatashare.DataSetKindSynapseWorkspaceSQLPoolTable),
	// 		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetProperties{
	// 			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
	// 		},
	// 	},
	// 	                        }
}
