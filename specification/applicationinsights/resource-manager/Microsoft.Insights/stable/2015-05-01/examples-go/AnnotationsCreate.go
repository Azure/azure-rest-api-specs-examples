package armapplicationinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsCreate.json
func ExampleAnnotationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAnnotationsClient().Create(ctx, "my-resource-group", "my-component", armapplicationinsights.Annotation{
		AnnotationName: to.Ptr("TestAnnotation"),
		Category:       to.Ptr("Text"),
		EventTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-31T13:41:38.657Z"); return t }()),
		ID:             to.Ptr("444e2c08-274a-4bbb-a89e-d77bb720f44a"),
		Properties:     to.Ptr("{\"Comments\":\"Testing\",\"Label\":\"Success\"}"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AnnotationArray = []*armapplicationinsights.Annotation{
	// 	{
	// 		AnnotationName: to.Ptr("TestAnnotation"),
	// 		Category: to.Ptr("Text"),
	// 		EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-31T13:41:38.657Z"); return t}()),
	// 		ID: to.Ptr("444e2c08-274a-4bbb-a89e-d77bb720f44a"),
	// 		Properties: to.Ptr("{\"Comments\":\"Testing\",\"Label\":\"Success\"}"),
	// }}
}
