package armintegrationspaces_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/azureintegrationspaces/resource-manager/Microsoft.IntegrationSpaces/preview/2023-11-14-preview/examples/Applications_ListBusinessProcessDevelopmentArtifacts.json
func ExampleApplicationsClient_ListBusinessProcessDevelopmentArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armintegrationspaces.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().ListBusinessProcessDevelopmentArtifacts(ctx, "testrg", "Space1", "Application1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListBusinessProcessDevelopmentArtifactsResponse = armintegrationspaces.ListBusinessProcessDevelopmentArtifactsResponse{
	// 	Value: []*armintegrationspaces.SaveOrGetBusinessProcessDevelopmentArtifactResponse{
	// 		{
	// 			Name: to.Ptr("BusinessProcess1"),
	// 			Properties: &armintegrationspaces.BusinessProcessDevelopmentArtifactProperties{
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
	// 								TrackingProfiles: map[string]*armintegrationspaces.TrackingProfileDefinition{
	// 									"subscriptions/sub1/resourcegroups/group1/providers/Microsoft.Web/sites/logicApp1": &armintegrationspaces.TrackingProfileDefinition{
	// 										Schema: to.Ptr("https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2023-01-01/trackingdefinitionschema.json#"),
	// 										BusinessProcess: &armintegrationspaces.BusinessProcessReference{
	// 											Name: to.Ptr("businessProcess1"),
	// 											Version: to.Ptr("d52c9c91-6e10-4a90-9c1f-08ee5d01c656"),
	// 										},
	// 										TrackingDefinitions: map[string]*armintegrationspaces.FlowTrackingDefinition{
	// 											"Fulfillment": &armintegrationspaces.FlowTrackingDefinition{
	// 												CorrelationContext: &armintegrationspaces.TrackingCorrelationContext{
	// 													OperationName: to.Ptr("manual"),
	// 													OperationType: to.Ptr("Trigger"),
	// 													PropertyName: to.Ptr("OrderNumber"),
	// 													Value: to.Ptr("@trigger().outputs.body.OrderNumber"),
	// 												},
	// 												Events: map[string]*armintegrationspaces.TrackingEventDefinition{
	// 													"Completed": &armintegrationspaces.TrackingEventDefinition{
	// 														OperationName: to.Ptr("CompletedPO"),
	// 														OperationType: to.Ptr("Action"),
	// 														Properties: map[string]any{
	// 														},
	// 													},
	// 													"Denied": &armintegrationspaces.TrackingEventDefinition{
	// 														OperationName: to.Ptr("DeniedPO"),
	// 														OperationType: to.Ptr("Action"),
	// 														Properties: map[string]any{
	// 														},
	// 													},
	// 													"Shipped": &armintegrationspaces.TrackingEventDefinition{
	// 														OperationName: to.Ptr("ShippedPO"),
	// 														OperationType: to.Ptr("Action"),
	// 														Properties: map[string]any{
	// 															"ShipPriority": "@action().inputs.shipPriority",
	// 															"TrackingID": "@action().inputs.trackingID",
	// 														},
	// 													},
	// 												},
	// 											},
	// 											"PurchaseOrder": &armintegrationspaces.FlowTrackingDefinition{
	// 												CorrelationContext: &armintegrationspaces.TrackingCorrelationContext{
	// 													OperationName: to.Ptr("manual"),
	// 													OperationType: to.Ptr("Trigger"),
	// 													PropertyName: to.Ptr("OrderNumber"),
	// 													Value: to.Ptr("@trigger().outputs.body.OrderNumber"),
	// 												},
	// 												Events: map[string]*armintegrationspaces.TrackingEventDefinition{
	// 													"Processing": &armintegrationspaces.TrackingEventDefinition{
	// 														OperationName: to.Ptr("ApprovedPO"),
	// 														OperationType: to.Ptr("Action"),
	// 														Properties: map[string]any{
	// 															"ApprovalStatus": "@action().inputs.ApprovalStatus",
	// 															"ApproverName": "@action().inputs.ApproverName",
	// 															"POAmount": "@action().inputs.POamount",
	// 														},
	// 													},
	// 													"Received": &armintegrationspaces.TrackingEventDefinition{
	// 														OperationName: to.Ptr("manual"),
	// 														OperationType: to.Ptr("Trigger"),
	// 														Properties: map[string]any{
	// 															"City": "@trigger().outputs.body.Address.City",
	// 															"Product": "@trigger().outputs.body.Product",
	// 															"Quantity": "@trigger().outputs.body.Quantity",
	// 															"State": "@trigger().outputs.body.Address.State",
	// 														},
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							SystemData: &armintegrationspaces.SystemData{
	// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T22:53:01.849Z"); return t}()),
	// 							},
	// 					}},
	// 				}
}
