package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/Relationships_Get.json
func ExampleRelationshipsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRelationshipsClient().Get(ctx, "rgopenapi", "myHealthModel", "Ue-21-F3M12V3w-13x18F8H-7HOk--kq6tP-HB", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.RelationshipsClientGetResponse{
	// 	Relationship: &armcloudhealth.Relationship{
	// 		Properties: &armcloudhealth.RelationshipProperties{
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			DisplayName: to.Ptr("My relationship"),
	// 			ParentEntityName: to.Ptr("Entity1"),
	// 			ChildEntityName: to.Ptr("Entity2"),
	// 			Labels: map[string]*string{
	// 				"key9681": to.Ptr("foo"),
	// 			},
	// 			DiscoveredBy: to.Ptr("discoveryRule1"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/myHealthModel/relationships/Ue-21-F3M12V3w-13x18F8H-7HOk--kq6tP-HB"),
	// 		Name: to.Ptr("Ue-21-F3M12V3w-13x18F8H-7HOk--kq6tP-HB"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels/relationships"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("arz"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
	// 		},
	// 	},
	// }
}
