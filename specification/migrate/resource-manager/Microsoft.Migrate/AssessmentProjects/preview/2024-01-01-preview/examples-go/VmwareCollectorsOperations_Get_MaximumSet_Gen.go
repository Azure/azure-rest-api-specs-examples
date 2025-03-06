package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/VmwareCollectorsOperations_Get_MaximumSet_Gen.json
func ExampleVmwareCollectorsOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVmwareCollectorsOperationsClient().Get(ctx, "ayagrawRG", "app18700project", "Vmware2258collector", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VmwareCollector = armmigrationassessment.VmwareCollector{
	// 	Name: to.Ptr("Vmware2258collector"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/vmwarecollectors"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/vmwarecollectors/Vmware2258collector"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.CollectorPropertiesBaseWithAgent{
	// 		AgentProperties: &armmigrationassessment.CollectorAgentPropertiesBase{
	// 			ID: to.Ptr("fe243486-3318-41fa-aaba-c48b5df75308"),
	// 			LastHeartbeatUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-29T12:10:08.916Z"); return t}()),
	// 			SpnDetails: &armmigrationassessment.CollectorAgentSpnPropertiesBase{
	// 				ApplicationID: to.Ptr("82b3e452-c0e8-4662-8347-58282925ae84"),
	// 				Audience: to.Ptr("82b3e452-c0e8-4662-8347-58282925ae84"),
	// 				Authority: to.Ptr("https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 				ObjectID: to.Ptr("3fc89111-1405-4938-9214-37aa4739401d"),
	// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			},
	// 			Version: to.Ptr("1.0.8.383"),
	// 		},
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-28T13:41:17.556Z"); return t}()),
	// 		DiscoverySiteID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/Vmware2744site"),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-16T12:23:45.969Z"); return t}()),
	// 	},
	// }
}
