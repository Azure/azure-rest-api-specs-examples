package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/HealthValidations_StartValidation_MaximumSet_Gen.json
func ExampleHealthValidationsClient_BeginStartValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("A76F9850-996B-40B3-94D4-C98110A0EEC9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewHealthValidationsClient().BeginStartValidation(ctx, "rgWatcher", "testWatcher", "testHealthValidation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.HealthValidationsClientStartValidationResponse{
	// 	HealthValidation: &armdatabasewatcher.HealthValidation{
	// 		Properties: &armdatabasewatcher.HealthValidationProperties{
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-07T19:42:55.594Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-07T19:42:55.594Z"); return t}()),
	// 			Status: to.Ptr(armdatabasewatcher.ValidationStatusRunning),
	// 			Issues: []*armdatabasewatcher.ValidationIssue{
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/469DD77C-C8DB-47B7-B9E1-72D29F8C878Be/resourceGroups/rgWatcher/providers/microsoft.databasewatcher/watchers/testWatcher/healthValidations/testHealthValidation"),
	// 		Name: to.Ptr("testHealthValidation"),
	// 		Type: to.Ptr("microsoft.databasewatcher/watchers/healthValidations"),
	// 		SystemData: &armdatabasewatcher.SystemData{
	// 			CreatedBy: to.Ptr("enbpvlpqbwd"),
	// 			CreatedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mxp"),
	// 			LastModifiedByType: to.Ptr(armdatabasewatcher.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-25T15:38:47.092Z"); return t}()),
	// 		},
	// 	},
	// }
}
