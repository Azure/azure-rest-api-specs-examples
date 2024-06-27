package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_CreateOrUpdate.json
func ExampleScopingConfigurationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopingConfigurationClient().CreateOrUpdate(ctx, "testReportName", "default", armappcomplianceautomation.ScopingConfigurationResource{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopingConfigurationResource = armappcomplianceautomation.ScopingConfigurationResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppComplianceAutomation/reports/scopingConfigurations"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName/scopingConfigurations/default"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ScopingConfigurationProperties{
	// 		Answers: []*armappcomplianceautomation.ScopingAnswer{
	// 			{
	// 				Answers: []*string{
	// 					to.Ptr("Azure")},
	// 					QuestionID: to.Ptr("GEN20_hostingEnvironment"),
	// 				},
	// 				{
	// 					Answers: []*string{
	// 					},
	// 					QuestionID: to.Ptr("DHP_G07_customerDataProcess"),
	// 				},
	// 				{
	// 					Answers: []*string{
	// 					},
	// 					QuestionID: to.Ptr("Tier2InitSub_serviceCommunicate"),
	// 			}},
	// 		},
	// 	}
}
