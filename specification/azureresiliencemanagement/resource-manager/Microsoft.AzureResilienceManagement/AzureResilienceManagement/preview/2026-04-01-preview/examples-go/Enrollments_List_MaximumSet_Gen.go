package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/Enrollments_List_MaximumSet_Gen.json
func ExampleEnrollmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnrollmentsClient().NewListPager("MyResourceGroup", "myUsagePlan", nil)
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
		// page = armresiliencemanagement.EnrollmentsClientListResponse{
		// 	EnrollmentListResult: armresiliencemanagement.EnrollmentListResult{
		// 		Value: []*armresiliencemanagement.Enrollment{
		// 			{
		// 				Properties: &armresiliencemanagement.EnrollmentProperties{
		// 					ServiceGroupID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sg1"),
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.AzureResilienceManagement/usagePlans/myUsagePlan/enrollments/sg1-enrollment"),
		// 				Name: to.Ptr("sg1-enrollment"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/usagePlans/enrollments"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T10:00:00.000Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T10:00:00.000Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/MyResourceGroup/providers/Microsoft.AzureResilienceManagement/usagePlans/myUsagePlan/enrollments?api-version=2026-04-01-preview&$skipToken=abc"),
		// 	},
		// }
	}
}
