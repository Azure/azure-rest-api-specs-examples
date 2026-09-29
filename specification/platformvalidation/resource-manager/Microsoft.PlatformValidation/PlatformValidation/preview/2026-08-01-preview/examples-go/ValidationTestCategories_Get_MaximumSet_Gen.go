package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/ValidationTestCategories_Get_MaximumSet_Gen.json
func ExampleValidationTestCategoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewValidationTestCategoriesClient().Get(ctx, "linux-quality-validations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplatformvalidation.ValidationTestCategoriesClientGetResponse{
	// 	ValidationTestCategory: armplatformvalidation.ValidationTestCategory{
	// 		Properties: &armplatformvalidation.ValidationTestCategoryProperties{
	// 			DisplayName: to.Ptr("Linux Quality Validations"),
	// 			Description: to.Ptr("AzCertify Linux image-quality validations. Groups Linux quality checks selected through the test's per-run inputs."),
	// 			Audience: to.Ptr(armplatformvalidation.CatalogAudiencePublic),
	// 			ProvisioningState: to.Ptr(armplatformvalidation.ResourceProvisioningStateSucceeded),
	// 			ParentCategoryID: to.Ptr("image-certification"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PlatformValidation/validationTestCategories/linux-quality-validations"),
	// 		Name: to.Ptr("linux-quality-validations"),
	// 		Type: to.Ptr("Microsoft.PlatformValidation/validationTestCategories"),
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
