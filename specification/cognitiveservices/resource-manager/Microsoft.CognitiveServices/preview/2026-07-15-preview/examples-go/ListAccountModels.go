package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/ListAccountModels.json
func ExampleAccountsClient_NewListModelsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListModelsPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.AccountsClientListModelsResponse{
		// 	AccountModelListResult: armcognitiveservices.AccountModelListResult{
		// 		Value: []*armcognitiveservices.AccountModel{
		// 			{
		// 				Name: to.Ptr("ada.1"),
		// 				Format: to.Ptr("OpenAI"),
		// 				BaseModel: &armcognitiveservices.DeploymentModel{
		// 					Name: to.Ptr("ada"),
		// 					Format: to.Ptr("OpenAI"),
		// 					Version: to.Ptr("1"),
		// 				},
		// 				Capabilities: map[string]*string{
		// 					"completion": to.Ptr("true"),
		// 					"fineTune": to.Ptr("true"),
		// 					"inference": to.Ptr("false"),
		// 				},
		// 				Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 					DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusPlanned),
		// 					FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 					Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 				},
		// 				IsDefaultVersion: to.Ptr(false),
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusLegacy),
		// 				MaxCapacity: to.Ptr[int32](10),
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2021, time.October, 7, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2021, time.October, 7, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("Microsoft"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("1"),
		// 			},
		// 			{
		// 				Name: to.Ptr("dall-e-3"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Capabilities: map[string]*string{
		// 					"imageGenerations": to.Ptr("true"),
		// 					"inference": to.Ptr("true"),
		// 				},
		// 				Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 					DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 					Inference: to.Ptr("2025-06-30T00:00:00Z"),
		// 				},
		// 				IsDefaultVersion: to.Ptr(true),
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusGenerallyAvailable),
		// 				MaxCapacity: to.Ptr[int32](2),
		// 				ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/dall-e-3/versions/3.0"),
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2023, time.August, 11, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2023, time.August, 11, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("Microsoft"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("3.0"),
		// 			},
		// 			{
		// 				Name: to.Ptr("gpt-35-turbo"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Capabilities: map[string]*string{
		// 					"chatCompletion": to.Ptr("true"),
		// 					"completion": to.Ptr("true"),
		// 					"fineTune": to.Ptr("false"),
		// 					"scaleType": to.Ptr("Manual,Standard"),
		// 				},
		// 				Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 					DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusPlanned),
		// 					Inference: to.Ptr("2025-04-30T00:00:00Z"),
		// 				},
		// 				IsDefaultVersion: to.Ptr(false),
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusDeprecated),
		// 				MaxCapacity: to.Ptr[int32](9),
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2023, time.March, 9, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2023, time.July, 6, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("Microsoft"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("0301"),
		// 			},
		// 			{
		// 				Name: to.Ptr("gpt-4o"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Capabilities: map[string]*string{
		// 					"chat": to.Ptr("true"),
		// 					"completion": to.Ptr("true"),
		// 					"fineTune": to.Ptr("false"),
		// 					"inference": to.Ptr("true"),
		// 					"vision": to.Ptr("true"),
		// 				},
		// 				Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 					DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 					Inference: to.Ptr("2025-09-15T00:00:00Z"),
		// 				},
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusDeprecating),
		// 				MaxCapacity: to.Ptr[int32](50),
		// 				ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/gpt-4o/versions/2024-05-13"),
		// 				ReplacementConfig: &armcognitiveservices.ReplacementConfig{
		// 					AutoUpgradeStartDate: to.Ptr(time.Date(2025, time.March, 26, 7, 0, 0, 0, time.UTC)),
		// 					TargetModelName: to.Ptr("gpt-4.1"),
		// 					TargetModelVersion: to.Ptr("2025-04-14"),
		// 					UpgradeOnExpiryLeadTimeDays: to.Ptr[int32](7),
		// 				},
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2024, time.May, 13, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("Microsoft"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("2024-05-13"),
		// 			},
		// 			{
		// 				Name: to.Ptr("Llama-3.2-90B-Vision-Instruct"),
		// 				Format: to.Ptr("Meta"),
		// 				Capabilities: map[string]*string{
		// 					"chatCompletion": to.Ptr("true"),
		// 				},
		// 				Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 					DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 					Inference: to.Ptr("2099-12-31T00:00:00Z"),
		// 				},
		// 				IsDefaultVersion: to.Ptr(false),
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusStable),
		// 				MaxCapacity: to.Ptr[int32](3),
		// 				ModelCatalogAssetID: to.Ptr("azureml://registries/azureml-meta/models/Llama-3.2-90B-Vision-Instruct/versions/2"),
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.April, 16, 4, 45, 33, 936787300, time.UTC)),
		// 					LastModifiedBy: to.Ptr("MaaSModelConverter"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("2"),
		// 			},
		// 			{
		// 				Name: to.Ptr("gpt-4o"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Capabilities: map[string]*string{
		// 					"chat": to.Ptr("true"),
		// 					"completion": to.Ptr("true"),
		// 					"fineTune": to.Ptr("false"),
		// 					"functionCalling": to.Ptr("true"),
		// 					"inference": to.Ptr("true"),
		// 					"vision": to.Ptr("true"),
		// 				},
		// 				FinetuneCapabilities: map[string]*string{
		// 					"chat": to.Ptr("true"),
		// 					"completion": to.Ptr("true"),
		// 					"fineTune": to.Ptr("true"),
		// 					"inference": to.Ptr("true"),
		// 					"scaleType": to.Ptr("Manual"),
		// 				},
		// 				LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusGenerallyAvailable),
		// 				MaxCapacity: to.Ptr[int32](50),
		// 				ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/gpt-4o/versions/2024-08-06"),
		// 				SystemData: &armcognitiveservices.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2024, time.August, 6, 0, 0, 0, 0, time.UTC)),
		// 					CreatedBy: to.Ptr("Microsoft"),
		// 					CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC)),
		// 					LastModifiedBy: to.Ptr("Microsoft"),
		// 					LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				},
		// 				Version: to.Ptr("2024-08-06"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
