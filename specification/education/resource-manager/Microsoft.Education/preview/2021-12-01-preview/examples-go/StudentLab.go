package armeducation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9b91929c304f8fb44002267b6c98d9fb9dde014/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLab.json
func ExampleStudentLabsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeducation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStudentLabsClient().Get(ctx, "{studentLabName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StudentLabDetails = armeducation.StudentLabDetails{
	// 	Name: to.Ptr("{studentLabName}"),
	// 	Type: to.Ptr("Microsoft.Education/StudentLabs"),
	// 	ID: to.Ptr(" /providers/Microsoft.Education/studentLabs/{studentLabName}"),
	// 	Properties: &armeducation.StudentLabProperties{
	// 		Description: to.Ptr("example lab description"),
	// 		Budget: &armeducation.Amount{
	// 			Currency: to.Ptr("USD"),
	// 			Value: to.Ptr[float32](100),
	// 		},
	// 		DisplayName: to.Ptr("example lab name"),
	// 		EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T21:54:17.438Z"); return t}()),
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-09T21:54:17.438Z"); return t}()),
	// 		LabScope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default"),
	// 		Role: to.Ptr(armeducation.StudentRoleStudent),
	// 		Status: to.Ptr(armeducation.StudentLabStatusActive),
	// 		SubscriptionID: to.Ptr("000000000-0000-0000-0000-00000000000000"),
	// 	},
	// }
}
