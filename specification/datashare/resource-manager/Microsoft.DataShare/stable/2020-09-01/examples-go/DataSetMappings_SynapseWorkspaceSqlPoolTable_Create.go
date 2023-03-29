package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/DataSetMappings_SynapseWorkspaceSqlPoolTable_Create.json
func ExampleDataSetMappingsClient_Create_dataSetMappingsSynapseWorkspaceSqlPoolTableCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataSetMappingsClient().Create(ctx, "SampleResourceGroup", "consumerAccount", "ShareSubscription1", "datasetMappingName1", &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMapping{
		Kind: to.Ptr(armdatashare.DataSetMappingKindSynapseWorkspaceSQLPoolTable),
		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMappingProperties{
			DataSetID:                              to.Ptr("3dc64e49-1fc3-4186-b3dc-d388c4d3076a"),
			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatashare.DataSetMappingsClientCreateResponse{
	// 	                            DataSetMappingClassification: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMapping{
	// 		Name: to.Ptr("datasetMappingName"),
	// 		Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions/dataSetMappings"),
	// 		ID: to.Ptr("/subscriptions/4e745bb7-c420-479b-b0d6-a0f92d48a227/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/consumerAccount/shareSubscriptions/ShareSubscription1/dataSetMappings/datasetMappingName1"),
	// 		Kind: to.Ptr(armdatashare.DataSetMappingKindSynapseWorkspaceSQLPoolTable),
	// 		Properties: &armdatashare.SynapseWorkspaceSQLPoolTableDataSetMappingProperties{
	// 			DataSetID: to.Ptr("3dc64e49-1fc3-4186-b3dc-d388c4d3076a"),
	// 			DataSetMappingStatus: to.Ptr(armdatashare.DataSetMappingStatusOk),
	// 			ProvisioningState: to.Ptr(armdatashare.ProvisioningStateSucceeded),
	// 			SynapseWorkspaceSQLPoolTableResourceID: to.Ptr("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"),
	// 		},
	// 	},
	// 	                        }
}
