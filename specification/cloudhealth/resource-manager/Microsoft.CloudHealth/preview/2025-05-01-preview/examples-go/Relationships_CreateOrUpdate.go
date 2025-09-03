package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/Relationships_CreateOrUpdate.json
func ExampleRelationshipsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRelationshipsClient().CreateOrUpdate(ctx, "rgopenapi", "model1", "rel1", armcloudhealth.Relationship{
		Properties: &armcloudhealth.RelationshipProperties{
			DisplayName:      to.Ptr("My relationship"),
			ParentEntityName: to.Ptr("Entity1"),
			ChildEntityName:  to.Ptr("Entity2"),
			Labels: map[string]*string{
				"key9681": to.Ptr("ixfvzsfnpvkkbrce"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.RelationshipsClientCreateOrUpdateResponse{
	// 	Relationship: &armcloudhealth.Relationship{
	// 		Properties: &armcloudhealth.RelationshipProperties{
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			DisplayName: to.Ptr("My relationship"),
	// 			ParentEntityName: to.Ptr("Entity1"),
	// 			ChildEntityName: to.Ptr("Entity2"),
	// 			Labels: map[string]*string{
	// 				"key9681": to.Ptr("ixfvzsfnpvkkbrce"),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/J1-j5J-E4N1r4I7i226K7-6V-B27V1RiF6S5M6-pl7JgD3-lx4CF/relationships/rel1"),
	// 		Name: to.Ptr("rel1"),
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
