package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/ValidationTests_Get_MaximumSet_Gen.json
func ExampleValidationTestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewValidationTestsClient().Get(ctx, "linux-quality-validation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplatformvalidation.ValidationTestsClientGetResponse{
	// 	ValidationTest: armplatformvalidation.ValidationTest{
	// 		Properties: &armplatformvalidation.ValidationTestProperties{
	// 			DisplayName: to.Ptr("Linux Quality Validation"),
	// 			Description: to.Ptr("Linux Quality Validation Suite backed by LISA. Runs the Linux quality checks selected by the caller against the image supplied in the execution plan."),
	// 			Audience: to.Ptr(armplatformvalidation.CatalogAudiencePublic),
	// 			ProvisioningState: to.Ptr(armplatformvalidation.ResourceProvisioningStateSucceeded),
	// 			CategoryIDs: []*string{
	// 				to.Ptr("linux-quality-validations"),
	// 			},
	// 			Inputs: []*armplatformvalidation.ValidationTestInput{
	// 				{
	// 					Name: to.Ptr("concurrency"),
	// 					Definition: &armplatformvalidation.ValidationTestInputDefinition{
	// 						Description: to.Ptr("Maximum number of Linux quality tests to run in parallel."),
	// 						Type: to.Ptr(armplatformvalidation.ValidationTestInputDataTypeInteger),
	// 						Required: to.Ptr(false),
	// 						DefaultValue: to.Ptr("0"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("testSuite"),
	// 					Definition: &armplatformvalidation.ValidationTestInputDefinition{
	// 						Description: to.Ptr("Linux quality test selections. Each array item can select checks using testNames, testArea, testCategory, testTags, or testPriority. The service runs the resolved union."),
	// 						Type: to.Ptr(armplatformvalidation.ValidationTestInputDataTypeArray),
	// 						Required: to.Ptr(true),
	// 					},
	// 				},
	// 			},
	// 			CurrentVersion: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation/versions/1.0.0"),
	// 			LatestPublishedVersion: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation/versions/1.0.0"),
	// 			LastPublishedAt: to.Ptr(time.Date(2026, time.September, 9, 11, 0, 0, 0, time.UTC)),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation"),
	// 		Name: to.Ptr("linux-quality-validation"),
	// 		Type: to.Ptr("Microsoft.PlatformValidation/validationTests"),
	// 		SystemData: &armplatformvalidation.SystemData{
	// 			CreatedBy: to.Ptr("11111111-1111-4111-8111-111111111111"),
	// 			CreatedByType: to.Ptr(armplatformvalidation.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.September, 9, 11, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("11111111-1111-4111-8111-111111111111"),
	// 			LastModifiedByType: to.Ptr(armplatformvalidation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.September, 9, 11, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
