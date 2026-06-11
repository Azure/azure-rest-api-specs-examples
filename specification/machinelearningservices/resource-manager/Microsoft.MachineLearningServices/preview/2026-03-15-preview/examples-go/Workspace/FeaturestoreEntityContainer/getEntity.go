package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/FeaturestoreEntityContainer/getEntity.json
func ExampleFeaturestoreEntityContainersClient_GetEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFeaturestoreEntityContainersClient().GetEntity(ctx, "test-rg", "my-aml-workspace", "string", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.FeaturestoreEntityContainersClientGetEntityResponse{
	// 	FeaturestoreEntityContainer: armmachinelearning.FeaturestoreEntityContainer{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Properties: &armmachinelearning.FeaturestoreEntityContainerProperties{
	// 			Description: to.Ptr("string"),
	// 			IsArchived: to.Ptr(false),
	// 			LatestVersion: to.Ptr("string"),
	// 			NextVersion: to.Ptr("string"),
	// 			Properties: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmachinelearning.AssetProvisioningStateUpdating),
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:43"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:43"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
