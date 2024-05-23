package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMax.json
func ExampleJobsClient_CreateOrUpdate_createAJobWithAllPropertiesSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().CreateOrUpdate(ctx, "group1", "server1", "agent1", "job1", armsql.Job{
		Properties: &armsql.JobProperties{
			Description: to.Ptr("my favourite job"),
			Schedule: &armsql.JobSchedule{
				Type:      to.Ptr(armsql.JobScheduleTypeRecurring),
				Enabled:   to.Ptr(true),
				EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59.000Z"); return t }()),
				Interval:  to.Ptr("PT5M"),
				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01.000Z"); return t }()),
			},
		},
	}, nil)
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
	// 		Description: to.Ptr("my favourite job"),
	// 		Schedule: &armsql.JobSchedule{
	// 			Type: to.Ptr(armsql.JobScheduleTypeRecurring),
	// 			Enabled: to.Ptr(true),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59.000Z"); return t}()),
	// 			Interval: to.Ptr("PT5M"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01.000Z"); return t}()),
	// 		},
	// 		Version: to.Ptr[int32](0),
	// 	},
	// }
}
