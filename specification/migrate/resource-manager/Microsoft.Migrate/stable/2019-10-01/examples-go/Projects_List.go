package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Projects_List.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListPager("abgoyal-westEurope", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProjectResultList = armmigrate.ProjectResultList{
		// 	Value: []*armmigrate.Project{
		// 		{
		// 			Name: to.Ptr("abgoyalWEselfhostb72bproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0600c777-0000-0d00-0000-5cdaa4170000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.migrate/migrateprojects/abgoyalweselfhost/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T08:28:53.330Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T11:18:47.789Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](3),
		// 				NumberOfGroups: to.Ptr[int32](2),
		// 				NumberOfMachines: to.Ptr[int32](28),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodwe.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:11:16.022Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
