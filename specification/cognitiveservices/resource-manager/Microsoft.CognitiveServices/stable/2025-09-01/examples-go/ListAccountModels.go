package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/ListAccountModels.json
func ExampleAccountsClient_NewListModelsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AccountModelListResult = armcognitiveservices.AccountModelListResult{
		// 	Value: []*armcognitiveservices.AccountModel{
		// 		{
		// 			Name: to.Ptr("ada.1"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("1"),
		// 			BaseModel: &armcognitiveservices.DeploymentModel{
		// 				Name: to.Ptr("ada"),
		// 				Format: to.Ptr("OpenAI"),
		// 				Version: to.Ptr("1"),
		// 			},
		// 			Capabilities: map[string]*string{
		// 				"completion": to.Ptr("true"),
		// 				"fineTune": to.Ptr("true"),
		// 				"inference": to.Ptr("false"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusPlanned),
		// 				FineTune: to.Ptr("2024-01-01T00:00:00Z"),
		// 				Inference: to.Ptr("2024-01-01T00:00:00Z"),
		// 			},
		// 			IsDefaultVersion: to.Ptr(false),
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusLegacy),
		// 			MaxCapacity: to.Ptr[int32](10),
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-07T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-07T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("Microsoft"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dall-e-3"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("3.0"),
		// 			Capabilities: map[string]*string{
		// 				"imageGenerations": to.Ptr("true"),
		// 				"inference": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 				Inference: to.Ptr("2025-06-30T00:00:00Z"),
		// 			},
		// 			IsDefaultVersion: to.Ptr(true),
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusGenerallyAvailable),
		// 			MaxCapacity: to.Ptr[int32](2),
		// 			ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/dall-e-3/versions/3.0"),
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-11T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-11T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("Microsoft"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("gpt-35-turbo"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("0301"),
		// 			Capabilities: map[string]*string{
		// 				"chatCompletion": to.Ptr("true"),
		// 				"completion": to.Ptr("true"),
		// 				"fineTune": to.Ptr("false"),
		// 				"scaleType": to.Ptr("Manual,Standard"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusPlanned),
		// 				Inference: to.Ptr("2025-04-30T00:00:00Z"),
		// 			},
		// 			IsDefaultVersion: to.Ptr(false),
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusDeprecated),
		// 			MaxCapacity: to.Ptr[int32](9),
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-09T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-06T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("Microsoft"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("gpt-4o"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("2024-05-13"),
		// 			Capabilities: map[string]*string{
		// 				"chat": to.Ptr("true"),
		// 				"completion": to.Ptr("true"),
		// 				"fineTune": to.Ptr("false"),
		// 				"inference": to.Ptr("true"),
		// 				"vision": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 				Inference: to.Ptr("2025-09-15T00:00:00Z"),
		// 			},
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusDeprecating),
		// 			MaxCapacity: to.Ptr[int32](50),
		// 			ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/gpt-4o/versions/2024-05-13"),
		// 			ReplacementConfig: &armcognitiveservices.ReplacementConfig{
		// 				AutoUpgradeStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-26T07:00:00.000Z"); return t}()),
		// 				TargetModelName: to.Ptr("gpt-4.1"),
		// 				TargetModelVersion: to.Ptr("2025-04-14"),
		// 				UpgradeOnExpiryLeadTimeDays: to.Ptr[int32](7),
		// 			},
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-13T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-15T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("Microsoft"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Llama-3.2-90B-Vision-Instruct"),
		// 			Format: to.Ptr("Meta"),
		// 			Version: to.Ptr("2"),
		// 			Capabilities: map[string]*string{
		// 				"chatCompletion": to.Ptr("true"),
		// 			},
		// 			Deprecation: &armcognitiveservices.ModelDeprecationInfo{
		// 				DeprecationStatus: to.Ptr(armcognitiveservices.DeprecationStatusTentative),
		// 				Inference: to.Ptr("2099-12-31T00:00:00Z"),
		// 			},
		// 			IsDefaultVersion: to.Ptr(false),
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusStable),
		// 			MaxCapacity: to.Ptr[int32](3),
		// 			ModelCatalogAssetID: to.Ptr("azureml://registries/azureml-meta/models/Llama-3.2-90B-Vision-Instruct/versions/2"),
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-16T04:45:33.936Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("MaaSModelConverter"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("gpt-4o"),
		// 			Format: to.Ptr("OpenAI"),
		// 			Version: to.Ptr("2024-08-06"),
		// 			Capabilities: map[string]*string{
		// 				"chat": to.Ptr("true"),
		// 				"completion": to.Ptr("true"),
		// 				"fineTune": to.Ptr("false"),
		// 				"functionCalling": to.Ptr("true"),
		// 				"inference": to.Ptr("true"),
		// 				"vision": to.Ptr("true"),
		// 			},
		// 			FinetuneCapabilities: map[string]*string{
		// 				"chat": to.Ptr("true"),
		// 				"completion": to.Ptr("true"),
		// 				"fineTune": to.Ptr("true"),
		// 				"inference": to.Ptr("true"),
		// 				"scaleType": to.Ptr("Manual"),
		// 			},
		// 			LifecycleStatus: to.Ptr(armcognitiveservices.ModelLifecycleStatusGenerallyAvailable),
		// 			MaxCapacity: to.Ptr[int32](50),
		// 			ModelCatalogAssetID: to.Ptr("azureml://registries/azure-openai/models/gpt-4o/versions/2024-08-06"),
		// 			SystemData: &armcognitiveservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-06T00:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("Microsoft"),
		// 				CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("Microsoft"),
		// 				LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
