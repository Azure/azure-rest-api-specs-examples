package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2026-05-01-preview/Relationships_CreateOrUpdate.json
func ExampleRelationshipsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("abcdef12-3456-7890-abcd-ef1234567890", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRelationshipsClient().BeginCreateOrUpdate(ctx, "online-store-rg", "online-store", "web-frontend-to-orders-api", armcloudhealth.Relationship{
		Properties: &armcloudhealth.RelationshipProperties{
			DisplayName:      to.Ptr("Web Frontend depends on Orders API"),
			ParentEntityName: to.Ptr("web-frontend"),
			ChildEntityName:  to.Ptr("orders-api"),
			Tags: map[string]*string{
				"environment": to.Ptr("production"),
				"team":        to.Ptr("online-store"),
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
	// res = armcloudhealth.RelationshipsClientCreateOrUpdateResponse{
	// 	Relationship: armcloudhealth.Relationship{
	// 		Properties: &armcloudhealth.RelationshipProperties{
	// 			DisplayName: to.Ptr("Web Frontend depends on Orders API"),
	// 			ParentEntityName: to.Ptr("web-frontend"),
	// 			ChildEntityName: to.Ptr("orders-api"),
	// 			Tags: map[string]*string{
	// 				"environment": to.Ptr("production"),
	// 				"team": to.Ptr("online-store"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/abcdef12-3456-7890-abcd-ef1234567890/resourceGroups/online-store-rg/providers/Microsoft.CloudHealth/healthmodels/online-store/relationships/web-frontend-to-orders-api"),
	// 		Name: to.Ptr("web-frontend-to-orders-api"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/relationships"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("admin@contoso.com"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T08:15:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("admin@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T09:30:00.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
