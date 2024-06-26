package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Transformation_Create.json
func ExampleTransformationsClient_CreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransformationsClient().CreateOrReplace(ctx, "sjrg4423", "sj8374", "transformation952", armstreamanalytics.Transformation{
		Properties: &armstreamanalytics.TransformationProperties{
			Query:          to.Ptr("Select Id, Name from inputtest"),
			StreamingUnits: to.Ptr[int32](6),
		},
	}, &armstreamanalytics.TransformationsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Transformation = armstreamanalytics.Transformation{
	// 	Name: to.Ptr("transformation952"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/transformations"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg4423/providers/Microsoft.StreamAnalytics/streamingjobs/sj8374/transformations/transformation952"),
	// 	Properties: &armstreamanalytics.TransformationProperties{
	// 		Query: to.Ptr("Select Id, Name from inputtest"),
	// 		StreamingUnits: to.Ptr[int32](6),
	// 	},
	// }
}
