package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Registry/ModelVersion/createOrUpdate.json
func ExampleRegistryModelVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistryModelVersionsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-registry", "string", "string", armmachinelearning.ModelVersion{
		Properties: &armmachinelearning.ModelVersionProperties{
			Description: to.Ptr("string"),
			Flavors: map[string]*armmachinelearning.FlavorData{
				"string": {
					Data: map[string]*string{
						"string": to.Ptr("string"),
					},
				},
			},
			IsAnonymous: to.Ptr(false),
			ModelType:   to.Ptr("CustomModel"),
			ModelURI:    to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.RegistryModelVersionsClientCreateOrUpdateResponse{
	// 	ModelVersion: armmachinelearning.ModelVersion{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Properties: &armmachinelearning.ModelVersionProperties{
	// 			Description: to.Ptr("string"),
	// 			Flavors: map[string]*armmachinelearning.FlavorData{
	// 				"string": &armmachinelearning.FlavorData{
	// 					Data: map[string]*string{
	// 						"string": to.Ptr("string"),
	// 					},
	// 				},
	// 			},
	// 			IsAnonymous: to.Ptr(false),
	// 			ModelType: to.Ptr("CustomModel"),
	// 			ModelURI: to.Ptr("string"),
	// 			Properties: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
