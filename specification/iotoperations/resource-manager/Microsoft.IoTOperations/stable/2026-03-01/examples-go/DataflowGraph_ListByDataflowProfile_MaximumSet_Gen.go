package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2026-03-01/DataflowGraph_ListByDataflowProfile_MaximumSet_Gen.json
func ExampleDataflowGraphClient_NewListByDataflowProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataflowGraphClient().NewListByDataflowProfilePager("rgiotoperations", "resource-123", "resource-123", nil)
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
		// page = armiotoperations.DataflowGraphClientListByDataflowProfileResponse{
		// 	DataflowGraphResourceListResult: armiotoperations.DataflowGraphResourceListResult{
		// 		Value: []*armiotoperations.DataflowGraphResource{
		// 			{
		// 				Properties: &armiotoperations.DataflowGraphProperties{
		// 					Mode: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 					RequestDiskPersistence: to.Ptr(armiotoperations.OperationalModeEnabled),
		// 					Nodes: []armiotoperations.DataflowGraphNodeClassification{
		// 						&armiotoperations.DataflowGraphSourceNode{
		// 							NodeType: to.Ptr(armiotoperations.DataflowGraphNodeTypeSource),
		// 							Name: to.Ptr("temperature"),
		// 							SourceSettings: &armiotoperations.DataflowGraphSourceSettings{
		// 								EndpointRef: to.Ptr("default"),
		// 								DataSources: []*string{
		// 									to.Ptr("telemetry/temperature"),
		// 								},
		// 							},
		// 						},
		// 						&armiotoperations.DataflowGraphGraphNode{
		// 							NodeType: to.Ptr(armiotoperations.DataflowGraphNodeTypeGraph),
		// 							Name: to.Ptr("my-graph"),
		// 							GraphSettings: &armiotoperations.DataflowGraphNodeGraphSettings{
		// 								RegistryEndpointRef: to.Ptr("my-registry-endpoint"),
		// 								Artifact: to.Ptr("my-wasm-module:1.4.3"),
		// 								Configuration: []*armiotoperations.DataflowGraphGraphNodeConfiguration{
		// 									{
		// 										Key: to.Ptr("key1"),
		// 										Value: to.Ptr("value1"),
		// 									},
		// 									{
		// 										Key: to.Ptr("key2"),
		// 										Value: to.Ptr("value2"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						&armiotoperations.DataflowGraphDestinationNode{
		// 							NodeType: to.Ptr(armiotoperations.DataflowGraphNodeTypeDestination),
		// 							Name: to.Ptr("alert"),
		// 							DestinationSettings: &armiotoperations.DataflowGraphDestinationNodeSettings{
		// 								EndpointRef: to.Ptr("default"),
		// 								DataDestination: to.Ptr("telemetry/temperature/alert"),
		// 							},
		// 						},
		// 						&armiotoperations.DataflowGraphDestinationNode{
		// 							NodeType: to.Ptr(armiotoperations.DataflowGraphNodeTypeDestination),
		// 							Name: to.Ptr("fabric"),
		// 							DestinationSettings: &armiotoperations.DataflowGraphDestinationNodeSettings{
		// 								EndpointRef: to.Ptr("fabric"),
		// 								DataDestination: to.Ptr("my-table"),
		// 							},
		// 						},
		// 					},
		// 					NodeConnections: []*armiotoperations.DataflowGraphNodeConnection{
		// 						{
		// 							From: &armiotoperations.DataflowGraphConnectionInput{
		// 								Name: to.Ptr("temperature"),
		// 								Schema: &armiotoperations.DataflowGraphConnectionSchemaSettings{
		// 									SerializationFormat: to.Ptr(armiotoperations.DataflowGraphConnectionSchemaSerializationFormatAvro),
		// 									SchemaRef: to.Ptr("aio-sr://namespace/temperature:1"),
		// 								},
		// 							},
		// 							To: &armiotoperations.DataflowGraphConnectionOutput{
		// 								Name: to.Ptr("my-graph"),
		// 							},
		// 						},
		// 						{
		// 							From: &armiotoperations.DataflowGraphConnectionInput{
		// 								Name: to.Ptr("my-graph.alert-output"),
		// 								Schema: &armiotoperations.DataflowGraphConnectionSchemaSettings{
		// 									SerializationFormat: to.Ptr(armiotoperations.DataflowGraphConnectionSchemaSerializationFormatAvro),
		// 									SchemaRef: to.Ptr("aio-sr://namespace/alert:1"),
		// 								},
		// 							},
		// 							To: &armiotoperations.DataflowGraphConnectionOutput{
		// 								Name: to.Ptr("fabric"),
		// 							},
		// 						},
		// 						{
		// 							From: &armiotoperations.DataflowGraphConnectionInput{
		// 								Name: to.Ptr("my-graph.normal-output"),
		// 								Schema: &armiotoperations.DataflowGraphConnectionSchemaSettings{
		// 									SerializationFormat: to.Ptr(armiotoperations.DataflowGraphConnectionSchemaSerializationFormatAvro),
		// 									SchemaRef: to.Ptr("aio-sr://namespace/alert:1"),
		// 								},
		// 							},
		// 							To: &armiotoperations.DataflowGraphConnectionOutput{
		// 								Name: to.Ptr("fabric"),
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.IoTOperations/instances/resource-123/dataflowProfiles/resource-123/dataflowGraphs/resource-123"),
		// 				ExtendedLocation: &armiotoperations.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123"),
		// 					Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
