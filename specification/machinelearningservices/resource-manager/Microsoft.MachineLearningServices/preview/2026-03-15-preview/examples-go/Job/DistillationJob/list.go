package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Job/DistillationJob/list.json
func ExampleJobsClient_NewListPager_listDistillationJob() {
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
		JobType:      to.Ptr("string"),
		ListViewType: to.Ptr(armmachinelearning.ListViewTypeAll),
		Tag:          to.Ptr("string")})
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
		// 				Properties: &armmachinelearning.DistillationJob{
		// 					Description: to.Ptr("string"),
		// 					ComponentID: to.Ptr("string"),
		// 					ComputeID: to.Ptr("string"),
		// 					DataGenerationDetails: &armmachinelearning.LabelGeneration{
		// 						DataGenerationTaskType: to.Ptr(armmachinelearning.DataGenerationTaskTypeConversation),
		// 						DataGenerationType: to.Ptr(armmachinelearning.DataGenerationTypeLabelGeneration),
		// 						TeacherModelEndpoint: &armmachinelearning.TeacherModelEndpoint{
		// 							EndpointName: to.Ptr("newfinetuneinttesting-jbuob"),
		// 						},
		// 						TrainingData: &armmachinelearning.URIFileJobInput{
		// 							Description: nil,
		// 							JobInputType: to.Ptr(armmachinelearning.JobInputTypeURIFile),
		// 							Mode: to.Ptr(armmachinelearning.InputDeliveryModeReadOnlyMount),
		// 							URI: to.Ptr("azureml://registries/azureml-meta/models/Llama-2-7b/versions/11"),
		// 						},
		// 					},
		// 					DisplayName: to.Ptr("string"),
		// 					ExperimentName: to.Ptr("string"),
		// 					FinetuningDetails: &armmachinelearning.FinetuningDetails{
		// 						StudentModel: &armmachinelearning.MLFlowModelJobInput{
		// 							Description: nil,
		// 							JobInputType: to.Ptr(armmachinelearning.JobInputTypeMlflowModel),
		// 							Mode: to.Ptr(armmachinelearning.InputDeliveryModeReadOnlyMount),
		// 							URI: to.Ptr("azureml://registries/azureml-meta/models/Meta-Llama-3.1-8B-Instruct/versions/1"),
		// 						},
		// 					},
		// 					Identity: &armmachinelearning.AmlToken{
		// 						IdentityType: to.Ptr(armmachinelearning.IdentityConfigurationTypeAMLToken),
		// 					},
		// 					IsArchived: to.Ptr(false),
		// 					JobType: to.Ptr(armmachinelearning.JobTypeDistillation),
		// 					NotificationSetting: &armmachinelearning.NotificationSetting{
		// 						EmailOn: []*armmachinelearning.EmailNotificationEnableType{
		// 							to.Ptr(armmachinelearning.EmailNotificationEnableTypeJobFailed),
		// 						},
		// 						Emails: []*string{
		// 							to.Ptr("string"),
		// 						},
		// 					},
		// 					Outputs: map[string]armmachinelearning.JobOutputClassification{
		// 						"string": &armmachinelearning.MLFlowModelJobOutput{
		// 							Description: to.Ptr("string"),
		// 							JobOutputType: to.Ptr(armmachinelearning.JobOutputTypeMlflowModel),
		// 							Mode: to.Ptr(armmachinelearning.OutputDeliveryModeReadWriteMount),
		// 							URI: to.Ptr("string"),
		// 						},
		// 					},
		// 					Properties: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 					QueueSettings: &armmachinelearning.QueueSettings{
		// 						JobTier: to.Ptr(armmachinelearning.JobTierStandard),
		// 					},
		// 					Resources: &armmachinelearning.JobResources{
		// 						InstanceTypes: []*string{
		// 							to.Ptr("Standard_NC6"),
		// 						},
		// 					},
		// 					Status: to.Ptr(armmachinelearning.JobStatus("Created")),
		// 					Tags: map[string]*string{
		// 						"string": to.Ptr("string"),
		// 					},
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T12:34:56.999+00:22"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T12:34:56.999+00:22"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
