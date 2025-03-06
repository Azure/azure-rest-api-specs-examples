package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AksSummaryOperations_Get_MaximumSet_Gen.json
func ExampleAksSummaryOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAksSummaryOperationsClient().Get(ctx, "rgaksswagger", "testproject", "testaksassessment", "AKS", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AKSSummary = armmigrationassessment.AKSSummary{
	// 	Name: to.Ptr("AKS"),
	// 	Type: to.Ptr("AKSAssessmentSummary"),
	// 	ID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentProjects/testproject/aksAssessments/testaksassessment/summaries/AKS"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
	// 		CreatedBy: to.Ptr("User"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	ETag: to.Ptr("00000000-0000-0000-a616-12d4724c01d9"),
	// 	Properties: &armmigrationassessment.AKSSummaryProperties{
	// 		AssessmentName: to.Ptr("testaksassessment"),
	// 		ConfidenceScore: to.Ptr[float32](7),
	// 		MonthlyComputeCost: to.Ptr[float32](8),
	// 		MonthlyStorageCost: to.Ptr[float32](26),
	// 		SuitabilityPerTarget: map[string]map[string]*int32{
	// 			"AKS": map[string]*int32{
	// 				"Ready": to.Ptr[int32](11),
	// 			},
	// 		},
	// 		SummaryName: to.Ptr("testaksassessmentsummary"),
	// 		AksAssessmentSummary: &armmigrationassessment.AKSAssessmentSummary{
	// 			MachineCount: to.Ptr[int32](17),
	// 			SKUSummary: map[string]*armmigrationassessment.SummaryDetails{
	// 				"StandardDS_v2": &armmigrationassessment.SummaryDetails{
	// 					Count: to.Ptr[int32](17),
	// 					MonthlyCost: to.Ptr[float32](18),
	// 				},
	// 			},
	// 			TotalMonthlyCost: to.Ptr[float32](6),
	// 			WebAppCount: to.Ptr[int32](19),
	// 			WebAppSummary: map[string]*armmigrationassessment.SummaryDetails{
	// 				"IIS": &armmigrationassessment.SummaryDetails{
	// 					Count: to.Ptr[int32](17),
	// 					MonthlyCost: to.Ptr[float32](18),
	// 				},
	// 			},
	// 			WebServerCount: to.Ptr[int32](28),
	// 			WebServerSummary: map[string]*int32{
	// 				"IIS": to.Ptr[int32](13),
	// 			},
	// 		},
	// 	},
	// }
}
