package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Entities_AddDataAnnotation.json
func ExampleEntitiesClient_AddDataAnnotation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().AddDataAnnotation(ctx, "online-store-rg", "online-store", "web-frontend", armcloudhealth.AddDataAnnotationRequest{
		AnnotationDetails: map[string]*string{
			"environment":  to.Ptr("production"),
			"deploymentId": to.Ptr("deploy-2026-05-04-001"),
			"changedBy":    to.Ptr("release-pipeline"),
		},
		Description: to.Ptr("Deployed release 2.4.1 to the web frontend."),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.EntitiesClientAddDataAnnotationResponse{
	// 	DataAnnotation: armcloudhealth.DataAnnotation{
	// 		AnnotationID: to.Ptr("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
	// 		AnnotationDetails: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"deploymentId": to.Ptr("deploy-2026-05-04-001"),
	// 			"changedBy": to.Ptr("release-pipeline"),
	// 		},
	// 		Description: to.Ptr("Deployed release 2.4.1 to the web frontend."),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T14:30:00Z"); return t}()),
	// 	},
	// }
}
