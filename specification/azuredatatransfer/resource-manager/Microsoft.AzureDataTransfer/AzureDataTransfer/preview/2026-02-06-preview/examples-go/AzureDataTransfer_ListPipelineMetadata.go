package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/AzureDataTransfer_ListPipelineMetadata.json
func ExampleAzureDataTransferClient_ListPipelineMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureDataTransferClient().ListPipelineMetadata(ctx, armdatatransfer.ListPipelineMetadataRequest{
		Pipeline: to.Ptr("testPipeline"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatatransfer.AzureDataTransferClientListPipelineMetadataResponse{
	// 	ListPipelineMetadataResult: armdatatransfer.ListPipelineMetadataResult{
	// 		Send: &armdatatransfer.DirectionSettings{
	// 			CompleteLinking: to.Ptr(armdatatransfer.EnablementStatusDisabled),
	// 			ConnectionApprovals: to.Ptr(armdatatransfer.EnablementStatusDisabled),
	// 			FlowTypesAvailable: []*armdatatransfer.FlowType{
	// 				to.Ptr(armdatatransfer.FlowTypeData),
	// 				to.Ptr(armdatatransfer.FlowTypeMessaging),
	// 				to.Ptr(armdatatransfer.FlowTypeMission),
	// 			},
	// 			MessagingFlowTypeSchemaValidation: to.Ptr(armdatatransfer.EnablementStatusDisabled),
	// 			FlowProfileDataClassesAvailable: []*armdatatransfer.DataClassType{
	// 				to.Ptr(armdatatransfer.DataClassTypeFiles),
	// 				to.Ptr(armdatatransfer.DataClassTypeMessaging),
	// 			},
	// 		},
	// 		Receive: &armdatatransfer.DirectionSettings{
	// 			CompleteLinking: to.Ptr(armdatatransfer.EnablementStatusEnabled),
	// 			ConnectionApprovals: to.Ptr(armdatatransfer.EnablementStatusEnabled),
	// 			FlowTypesAvailable: []*armdatatransfer.FlowType{
	// 				to.Ptr(armdatatransfer.FlowTypeData),
	// 				to.Ptr(armdatatransfer.FlowTypeMessaging),
	// 				to.Ptr(armdatatransfer.FlowTypeMission),
	// 			},
	// 			MessagingFlowTypeSchemaValidation: to.Ptr(armdatatransfer.EnablementStatusDisabled),
	// 			FlowProfileDataClassesAvailable: []*armdatatransfer.DataClassType{
	// 				to.Ptr(armdatatransfer.DataClassTypeFiles),
	// 				to.Ptr(armdatatransfer.DataClassTypeMessaging),
	// 			},
	// 		},
	// 	},
	// }
}
