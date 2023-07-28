package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMin.json
func ExampleJobsClient_CreateOrUpdate_createAJobWithDefaultProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "group1", "server1", "agent1", "job1", armsql.Job{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armsql.Job{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAccounts/jobs"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/jobs/job1"),
	// 	Properties: &armsql.JobProperties{
	// 		Description: to.Ptr(""),
	// 		Schedule: &armsql.JobSchedule{
	// 			Type: to.Ptr(armsql.JobScheduleTypeOnce),
	// 			Enabled: to.Ptr(false),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T11:59:59Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z"); return t}()),
	// 		},
	// 		Version: to.Ptr[int32](0),
	// 	},
	// }
}
