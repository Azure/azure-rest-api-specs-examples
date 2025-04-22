package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Delete.json
func ExampleDataFlowDebugSessionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataFlowDebugSessionClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DeleteDataFlowDebugSessionRequest{
		SessionID: to.Ptr("91fb57e0-8292-47be-89ff-c8f2d2bb2a7e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
