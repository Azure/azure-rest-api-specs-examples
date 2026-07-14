package armcloudhealth_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_GetDataAnnotations.json
func ExampleEntitiesClient_GetDataAnnotations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().GetDataAnnotations(ctx, "online-store-rg", "online-store", "web-frontend", armcloudhealth.GetDataAnnotationsRequest{
		StartAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-03T00:00:00Z"); return t }()),
		EndAt:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T23:59:59Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientGetDataAnnotationsResponse{
	// 	GetDataAnnotationsResponse: armcloudhealth.GetDataAnnotationsResponse{
	// 		EntityName: to.Ptr("web-frontend"),
	// 		Annotations: []*armcloudhealth.DataAnnotation{
	// 			{
	// 				AnnotationID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T14:30:00Z"); return t}()),
	// 				AnnotationDetails: map[string]*string{
	// 					"environment": to.Ptr("production"),
	// 					"deploymentId": to.Ptr("deploy-2026-05-04-001"),
	// 					"changedBy": to.Ptr("release-pipeline"),
	// 				},
	// 				Description: to.Ptr("Deployed release 2.4.1 to the web frontend."),
	// 			},
	// 			{
	// 				AnnotationID: to.Ptr("b2c3d4e5-f6a7-8901-bcde-f21234567890"),
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T03:15:00Z"); return t}()),
	// 				AnnotationDetails: map[string]*string{
	// 					"changeType": to.Ptr("ScaleOut"),
	// 					"instanceCount": to.Ptr("6"),
	// 				},
	// 				Description: to.Ptr("Scaled out the web frontend to 6 instances."),
	// 			},
	// 		},
	// 	},
	// }
}
