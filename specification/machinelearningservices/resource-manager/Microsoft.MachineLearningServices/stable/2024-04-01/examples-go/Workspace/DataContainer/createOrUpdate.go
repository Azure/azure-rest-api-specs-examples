package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/DataContainer/createOrUpdate.json
func ExampleDataContainersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataContainersClient().CreateOrUpdate(ctx, "testrg123", "workspace123", "datacontainer123", armmachinelearning.DataContainer{
		Properties: &armmachinelearning.DataContainerProperties{
			Description: to.Ptr("string"),
			Properties: map[string]*string{
				"properties1": to.Ptr("value1"),
				"properties2": to.Ptr("value2"),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			DataType: to.Ptr(armmachinelearning.DataType("UriFile")),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataContainer = armmachinelearning.DataContainer{
	// 	Name: to.Ptr("datacontainer123"),
	// 	Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/data"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspace123/data/datacontainer123"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("John Smith"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("John Smith"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmachinelearning.DataContainerProperties{
	// 		Description: to.Ptr("string"),
	// 		Properties: map[string]*string{
	// 			"properties1": to.Ptr("value1"),
	// 			"properties2": to.Ptr("value2"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 		DataType: to.Ptr(armmachinelearning.DataType("UriFile")),
	// 	},
	// }
}
