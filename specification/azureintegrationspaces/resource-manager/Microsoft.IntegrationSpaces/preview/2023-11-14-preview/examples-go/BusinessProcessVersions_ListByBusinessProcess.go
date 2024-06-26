package armintegrationspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/azureintegrationspaces/resource-manager/Microsoft.IntegrationSpaces/preview/2023-11-14-preview/examples/BusinessProcessVersions_ListByBusinessProcess.json
func ExampleBusinessProcessVersionsClient_NewListByBusinessProcessPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armintegrationspaces.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBusinessProcessVersionsClient().NewListByBusinessProcessPager("testrg", "Space1", "Application1", "BusinessProcess1", &armintegrationspaces.BusinessProcessVersionsClientListByBusinessProcessOptions{Top: nil,
		Skip:        nil,
		Maxpagesize: nil,
		Filter:      nil,
		Select:      []string{},
		Expand:      []string{},
		Orderby:     []string{},
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BusinessProcessVersionListResult = armintegrationspaces.BusinessProcessVersionListResult{
		// 	Value: []*armintegrationspaces.BusinessProcessVersion{
		// 		{
		// 			Name: to.Ptr("08585074782265427079"),
		// 			Type: to.Ptr("Microsoft.IntegrationSpaces/spaces/applications/BusinessProcesses/versions"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.IntegrationSpaces/spaces/Space1/applications/Application1/businessProcesses/BusinessProcess1/versions/08585074782265427079"),
		// 			Properties: &armintegrationspaces.BusinessProcessProperties{
		// 				Description: to.Ptr("First Business Process"),
		// 				BusinessProcessMapping: map[string]*armintegrationspaces.BusinessProcessMappingItem{
		// 					"Completed": &armintegrationspaces.BusinessProcessMappingItem{
		// 						LogicAppResourceID: to.Ptr("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
		// 						OperationName: to.Ptr("CompletedPO"),
		// 						OperationType: to.Ptr("Action"),
		// 						WorkflowName: to.Ptr("Fulfillment"),
		// 					},
		// 					"Denied": &armintegrationspaces.BusinessProcessMappingItem{
		// 						LogicAppResourceID: to.Ptr("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
		// 						OperationName: to.Ptr("DeniedPO"),
		// 						OperationType: to.Ptr("Action"),
		// 						WorkflowName: to.Ptr("Fulfillment"),
		// 					},
		// 					"Processing": &armintegrationspaces.BusinessProcessMappingItem{
		// 						LogicAppResourceID: to.Ptr("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
		// 						OperationName: to.Ptr("ApprovedPO"),
		// 						OperationType: to.Ptr("Action"),
		// 						WorkflowName: to.Ptr("PurchaseOrder"),
		// 					},
		// 					"Received": &armintegrationspaces.BusinessProcessMappingItem{
		// 						LogicAppResourceID: to.Ptr("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
		// 						OperationName: to.Ptr("manual"),
		// 						OperationType: to.Ptr("Trigger"),
		// 						WorkflowName: to.Ptr("PurchaseOrder"),
		// 					},
		// 					"Shipped": &armintegrationspaces.BusinessProcessMappingItem{
		// 						LogicAppResourceID: to.Ptr("subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1"),
		// 						OperationName: to.Ptr("ShippedPO"),
		// 						OperationType: to.Ptr("Action"),
		// 						WorkflowName: to.Ptr("Fulfillment"),
		// 					},
		// 				},
		// 				BusinessProcessStages: map[string]*armintegrationspaces.BusinessProcessStage{
		// 					"Completed": &armintegrationspaces.BusinessProcessStage{
		// 						Description: to.Ptr("Completed"),
		// 						StagesBefore: []*string{
		// 							to.Ptr("Shipped")},
		// 						},
		// 						"Denied": &armintegrationspaces.BusinessProcessStage{
		// 							Description: to.Ptr("Denied"),
		// 							StagesBefore: []*string{
		// 								to.Ptr("Processing")},
		// 							},
		// 							"Processing": &armintegrationspaces.BusinessProcessStage{
		// 								Description: to.Ptr("Processing"),
		// 								Properties: map[string]*string{
		// 									"ApprovalState": to.Ptr("String"),
		// 									"ApproverName": to.Ptr("String"),
		// 									"POAmount": to.Ptr("Integer"),
		// 								},
		// 								StagesBefore: []*string{
		// 									to.Ptr("Received")},
		// 								},
		// 								"Received": &armintegrationspaces.BusinessProcessStage{
		// 									Description: to.Ptr("received"),
		// 									Properties: map[string]*string{
		// 										"City": to.Ptr("String"),
		// 										"Product": to.Ptr("String"),
		// 										"Quantity": to.Ptr("Integer"),
		// 										"State": to.Ptr("String"),
		// 									},
		// 								},
		// 								"Shipped": &armintegrationspaces.BusinessProcessStage{
		// 									Description: to.Ptr("Shipped"),
		// 									Properties: map[string]*string{
		// 										"ShipPriority": to.Ptr("Integer"),
		// 										"TrackingID": to.Ptr("Integer"),
		// 									},
		// 									StagesBefore: []*string{
		// 										to.Ptr("Denied")},
		// 									},
		// 								},
		// 								Identifier: &armintegrationspaces.BusinessProcessIdentifier{
		// 									PropertyName: to.Ptr("businessIdentifier-1"),
		// 									PropertyType: to.Ptr("String"),
		// 								},
		// 								ProvisioningState: to.Ptr(armintegrationspaces.ProvisioningStateSucceeded),
		// 								TableName: to.Ptr("table1"),
		// 								TrackingDataStoreReferenceName: to.Ptr("trackingDataStoreReferenceName1"),
		// 								Version: to.Ptr("08585074782265427079"),
		// 							},
		// 					}},
		// 				}
	}
}
