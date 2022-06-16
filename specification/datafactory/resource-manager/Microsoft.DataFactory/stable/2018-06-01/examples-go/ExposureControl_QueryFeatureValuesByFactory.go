package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_QueryFeatureValuesByFactory.json
func ExampleExposureControlClient_QueryFeatureValuesByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewExposureControlClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.QueryFeatureValuesByFactory(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		armdatafactory.ExposureControlBatchRequest{
			ExposureControlRequests: []*armdatafactory.ExposureControlRequest{
				{
					FeatureName: to.Ptr("ADFIntegrationRuntimeSharingRbac"),
					FeatureType: to.Ptr("Feature"),
				},
				{
					FeatureName: to.Ptr("ADFSampleFeature"),
					FeatureType: to.Ptr("Feature"),
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
