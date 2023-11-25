package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Groups_Get.json
func ExampleGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGroupsClient().Get(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "Test1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Group = armmigrate.Group{
	// 	Name: to.Ptr("Test1"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups"),
	// 	ETag: to.Ptr("\"21009c31-0000-0d00-0000-5cd585ad0000\""),
	// 	ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/Test1"),
	// 	Properties: &armmigrate.GroupProperties{
	// 		AreAssessmentsRunning: to.Ptr(false),
	// 		Assessments: []*string{
	// 			to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/Test1/assessments/assessment_5_9_2019_16_22_14")},
	// 			CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:52:07.736Z"); return t}()),
	// 			GroupStatus: to.Ptr(armmigrate.GroupStatusCompleted),
	// 			MachineCount: to.Ptr[int32](26),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-10T14:07:41.752Z"); return t}()),
	// 		},
	// 	}
}
