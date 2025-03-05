package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/EvaluatedSqlEntitiesOperations_Get_MaximumSet_Gen.json
func ExampleEvaluatedSQLEntitiesOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEvaluatedSQLEntitiesOperationsClient().Get(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", "a404-r1w16-1.FPL.COM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EvaluatedSQLEntity = armmigrationassessment.EvaluatedSQLEntity{
	// 	Name: to.Ptr("a404-r1w16-1.FPL.COM"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/businessCases/evaluatedsqlentity"),
	// 	ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentprojects/multipleto8617project/businessCases/sample-business-case/evaluatedsqlentity/a404-r1w16-1.FPL.COM"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		CreatedBy: to.Ptr("t72jdt@company.com"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("t72jdt@company.com"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.EvaluatedSQLEntityProperties{
	// 		ActivityState: to.Ptr("active"),
	// 		DbCount: to.Ptr[int32](1),
	// 		QualifyingOffer: to.Ptr("3 Year RI"),
	// 		ReadyForMigration: to.Ptr("Ready"),
	// 		RecommendedAzureTarget: to.Ptr("AzureVirtualMachine"),
	// 		ServerName: to.Ptr("a404-r1w16-1.FPL.COM"),
	// 		SupportStatus: to.Ptr(armmigrationassessment.SupportabilityStatusUnknown),
	// 		VirtualizationType: to.Ptr("Vmware"),
	// 	},
	// }
}
