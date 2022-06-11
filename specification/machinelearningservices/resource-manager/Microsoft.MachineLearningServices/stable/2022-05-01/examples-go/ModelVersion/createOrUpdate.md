```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/ModelVersion/createOrUpdate.json
func ExampleModelVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewModelVersionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"string",
		"string",
		armmachinelearning.ModelVersion{
			Properties: &armmachinelearning.ModelVersionProperties{
				Description: to.Ptr("string"),
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				Tags: map[string]*string{
					"string": to.Ptr("string"),
				},
				IsAnonymous: to.Ptr(false),
				Flavors: map[string]*armmachinelearning.FlavorData{
					"string": {
						Data: map[string]*string{
							"string": to.Ptr("string"),
						},
					},
				},
				ModelType: to.Ptr("CustomModel"),
				ModelURI:  to.Ptr("string"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
