package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/PrivateLinkResources_ListByProject.json
func ExamplePrivateLinkResourceClient_ListByProject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourceClient().ListByProject(ctx, "madhavicus", "custestpece80project", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceCollection = armmigrate.PrivateLinkResourceCollection{
	// 	Value: []*armmigrate.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("Default"),
	// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/madhavicus/providers/Microsoft.Migrate/assessmentprojects/custestpece80project/privateLinkResources/Default"),
	// 			Properties: &armmigrate.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("Default"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("CollectorAgent")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.prod.migration.windowsazure.com")},
	// 					},
	// 			}},
	// 		}
}
