package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/ModelVersion/list.json
func ExampleModelVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewModelVersionsClient().NewListPager("test-rg", "my-aml-workspace", "string", &armmachinelearning.ModelVersionsClientListOptions{
		Description: to.Ptr("string"),
		OrderBy:     to.Ptr("string"),
		Top:         to.Ptr[int32](1),
		Offset:      to.Ptr[int32](1),
		Properties:  to.Ptr("string"),
		Tags:        to.Ptr("string"),
		Version:     to.Ptr("string")})
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
		// page = armmachinelearning.ModelVersionsClientListResponse{
		// 	ModelVersionResourceArmPaginatedResult: armmachinelearning.ModelVersionResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/models/string/versions?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.ModelVersion{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Properties: &armmachinelearning.ModelVersionProperties{
		// 					Description: to.Ptr("string"),
		// 					Flavors: map[string]*armmachinelearning.FlavorData{
		// 						"string": &armmachinelearning.FlavorData{
		// 							Data: map[string]*string{
		// 								"string": to.Ptr("string"),
		// 							},
		// 						},
		// 					},
		// 					IsAnonymous: to.Ptr(false),
		// 					ModelType: to.Ptr("CustomModel"),
		// 					ModelURI: to.Ptr("string"),
		// 					Properties: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					Tags: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
