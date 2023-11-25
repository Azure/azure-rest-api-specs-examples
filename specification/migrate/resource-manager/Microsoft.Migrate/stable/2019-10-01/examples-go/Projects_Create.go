package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Projects_Create.json
func ExampleProjectsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Create(ctx, "abgoyal-westEurope", "abGoyalProject2", &armmigrate.ProjectsClientCreateOptions{Project: &armmigrate.Project{
		ETag:     to.Ptr(""),
		Location: to.Ptr("West Europe"),
		Properties: &armmigrate.ProjectProperties{
			AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.migrate/migrateprojects/abgoyalweselfhost/Solutions/Servers-Assessment-ServerAssessment"),
			ProjectStatus:        to.Ptr(armmigrate.ProjectStatusActive),
		},
		Tags: map[string]any{},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armmigrate.Project{
	// 	Name: to.Ptr("abGoyalProject2"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
	// 	ETag: to.Ptr(""),
	// 	ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abGoyalProject2"),
	// 	Location: to.Ptr("West Europe"),
	// 	Properties: &armmigrate.ProjectProperties{
	// 		AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.Migrate/assessmentprojects/abGoyalProject2"),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-26T05:56:58.752Z"); return t}()),
	// 		NumberOfAssessments: to.Ptr[int32](0),
	// 		NumberOfGroups: to.Ptr[int32](0),
	// 		NumberOfMachines: to.Ptr[int32](0),
	// 		ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
	// 		ServiceEndpoint: to.Ptr("https://localhost/"),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-26T05:56:58.799Z"); return t}()),
	// 	},
	// 	Tags: map[string]any{
	// 	},
	// }
}
