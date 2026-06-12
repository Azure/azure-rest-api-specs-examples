package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/DataContainer/list.json
func ExampleDataContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataContainersClient().NewListPager("testrg123", "workspace123", nil)
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
		// page = armmachinelearning.DataContainersClientListResponse{
		// 	DataContainerResourceArmPaginatedResult: armmachinelearning.DataContainerResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspaces123/data?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.DataContainer{
		// 			{
		// 				Name: to.Ptr("datastore123"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/data"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspace123/data/datacontainer123"),
		// 				Properties: &armmachinelearning.DataContainerProperties{
		// 					Description: to.Ptr("string"),
		// 					DataType: to.Ptr(armmachinelearning.DataType("UriFile")),
		// 					Properties: map[string]*string{
		// 						"properties1": to.Ptr("value1"),
		// 						"properties2": to.Ptr("value2"),
		// 					},
		// 					Tags: map[string]*string{
		// 						"tag1": to.Ptr("value1"),
		// 						"tag2": to.Ptr("value2"),
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 					CreatedBy: to.Ptr("John Smith"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("John Smith"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("datastore124"),
		// 				Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/data"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/workspace123/data/datacontainer124"),
		// 				Properties: &armmachinelearning.DataContainerProperties{
		// 					Description: to.Ptr("string"),
		// 					DataType: to.Ptr(armmachinelearning.DataType("UriFile")),
		// 					Properties: map[string]*string{
		// 						"properties1": to.Ptr("value1"),
		// 						"properties2": to.Ptr("value2"),
		// 					},
		// 					Tags: map[string]*string{
		// 						"tag1": to.Ptr("value1"),
		// 						"tag2": to.Ptr("value2"),
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 					CreatedBy: to.Ptr("John Smith"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T12:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("John Smith"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
