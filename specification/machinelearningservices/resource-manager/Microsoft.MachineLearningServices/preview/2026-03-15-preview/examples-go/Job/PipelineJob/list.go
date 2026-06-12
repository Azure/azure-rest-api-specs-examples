package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Job/PipelineJob/list.json
func ExampleJobsClient_NewListPager_listPipelineJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearning.JobsClientListOptions{
		JobType: to.Ptr("string"),
		Tag:     to.Ptr("string")})
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
		// page = armmachinelearning.JobsClientListResponse{
		// 	JobBaseResourceArmPaginatedResult: armmachinelearning.JobBaseResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/jobs?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.JobBase{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Properties: &armmachinelearning.PipelineJob{
		// 					Description: to.Ptr("string"),
		// 					ComputeID: to.Ptr("string"),
		// 					DisplayName: to.Ptr("string"),
		// 					ExperimentName: to.Ptr("string"),
		// 					Inputs: map[string]armmachinelearning.JobInputClassification{
		// 						"string": &armmachinelearning.LiteralJobInput{
		// 							Description: to.Ptr("string"),
		// 							JobInputType: to.Ptr(armmachinelearning.JobInputTypeLiteral),
		// 							Value: to.Ptr("string"),
		// 						},
		// 					},
		// 					JobType: to.Ptr(armmachinelearning.JobTypePipeline),
		// 					Outputs: map[string]armmachinelearning.JobOutputClassification{
		// 						"string": &armmachinelearning.URIFileJobOutput{
		// 							Description: to.Ptr("string"),
		// 							JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeURIFile),
		// 							Mode: to.Ptr(armmachinelearning.OutputDeliveryModeUpload),
		// 							URI: to.Ptr("string"),
		// 						},
		// 					},
		// 					Properties: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					Services: map[string]*armmachinelearning.JobService{
		// 						"string": &armmachinelearning.JobService{
		// 							Endpoint: to.Ptr("string"),
		// 							ErrorMessage: to.Ptr("string"),
		// 							JobServiceType: to.Ptr("string"),
		// 							Port: to.Ptr[int32](1),
		// 							Properties: map[string]*string{
		// 								"string": to.Ptr("string"),
		// 							},
		// 							Status: to.Ptr("string"),
		// 						},
		// 					},
		// 					Settings: map[string]any{
		// 					},
		// 					Status: to.Ptr(armmachinelearning.JobStatusNotStarted),
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
