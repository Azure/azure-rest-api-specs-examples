package armrelationships_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relationships/armrelationships"
)

// Generated from example definition: 2026-08-01/DependencyOfRelationshipsByServiceGroup_Get.json
func ExampleDependencyOfRelationshipsByServiceGroupClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelationships.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDependencyOfRelationshipsByServiceGroupClient().Get(ctx, "myServiceGroup", "relationshipOne", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelationships.DependencyOfRelationshipsByServiceGroupClientGetResponse{
	// 	DependencyOfRelationship: armrelationships.DependencyOfRelationship{
	// 		Properties: &armrelationships.DependencyOfRelationshipProperties{
	// 			SourceID: to.Ptr("/providers/Microsoft.Management/serviceGroups/myServiceGroup"),
	// 			TargetID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg123/providers/Microsoft.Web/staticSites/test-site"),
	// 			TargetTenant: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			OriginInformation: &armrelationships.RelationshipOriginInformation{
	// 				RelationshipOriginType: to.Ptr(armrelationships.RelationshipOriginsUserExplicitlyCreated),
	// 				DiscoveryEngine: to.Ptr("PEM"),
	// 			},
	// 			Metadata: &armrelationships.RelationshipMetadata{
	// 				SourceType: to.Ptr("Microsoft.Management/serviceGroups"),
	// 				TargetType: to.Ptr("Microsoft.Web/staticSites"),
	// 			},
	// 			ProvisioningState: to.Ptr(armrelationships.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/myServiceGroup/providers/Microsoft.Relationships/dependencyOf/relationshipOne"),
	// 		Name: to.Ptr("relationshipOne"),
	// 		Type: to.Ptr("Microsoft.Relationships/dependencyOf"),
	// 	},
	// }
}
