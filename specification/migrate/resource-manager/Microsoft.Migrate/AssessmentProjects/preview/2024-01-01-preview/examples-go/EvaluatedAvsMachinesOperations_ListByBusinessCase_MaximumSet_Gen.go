package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/EvaluatedAvsMachinesOperations_ListByBusinessCase_MaximumSet_Gen.json
func ExampleEvaluatedAvsMachinesOperationsClient_NewListByBusinessCasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEvaluatedAvsMachinesOperationsClient().NewListByBusinessCasePager("rgopenapi", "multipleto8617project", "sample-business-case", &armmigrationassessment.EvaluatedAvsMachinesOperationsClientListByBusinessCaseOptions{Filter: to.Ptr("zcwvgkjkvddoylnfkgclpytp"),
		PageSize:          to.Ptr[int32](9),
		ContinuationToken: to.Ptr("er"),
		TotalRecordCount:  to.Ptr[int32](2),
	})
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
		// page.EvaluatedAvsMachineListResult = armmigrationassessment.EvaluatedAvsMachineListResult{
		// 	Value: []*armmigrationassessment.EvaluatedAvsMachine{
		// 		{
		// 			Name: to.Ptr("a404-r1w16-1.FPL.COM"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/businessCases/evaluatedavsmachine"),
		// 			ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentprojects/multipleto8617project/businessCases/sample-business-case/evaluatedavsmachine/a404-r1w16-1.FPL.COM"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
		// 				CreatedBy: to.Ptr("t72jdt@company.com"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("t72jdt@company.com"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.EvaluatedAvsMachineProperties{
		// 				ActivityState: to.Ptr("inconclusive"),
		// 				MachineID: to.Ptr("a404-r1w16-1.FPL.COM"),
		// 				OperatingSystemName: to.Ptr("Ubuntu"),
		// 				QualifyingOffer: to.Ptr("3 Year RI"),
		// 				ReadyForMigration: to.Ptr("Ready"),
		// 				RecommendedAzureTarget: to.Ptr("AzureVm"),
		// 				ServerName: to.Ptr("a404-r1w16-1.FPL.COM"),
		// 				SupportStatus: to.Ptr(armmigrationassessment.SupportabilityStatusUnknown),
		// 				VirtualizationType: to.Ptr("Vmware"),
		// 			},
		// 	}},
		// }
	}
}
