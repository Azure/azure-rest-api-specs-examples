package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_GetScopingQuestions.json
func ExampleReportClient_GetScopingQuestions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().GetScopingQuestions(ctx, "testReportName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopingQuestions = armappcomplianceautomation.ScopingQuestions{
	// 	Questions: []*armappcomplianceautomation.ScopingQuestion{
	// 		{
	// 			InputType: to.Ptr(armappcomplianceautomation.InputTypeBoolean),
	// 			OptionIDs: []*string{
	// 			},
	// 			QuestionID: to.Ptr("DHP_G07_customerDataProcess"),
	// 			Rules: []*armappcomplianceautomation.Rule{
	// 				to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 				ShowSubQuestionsValue: to.Ptr("true"),
	// 			},
	// 			{
	// 				InputType: to.Ptr(armappcomplianceautomation.InputTypeText),
	// 				OptionIDs: []*string{
	// 				},
	// 				QuestionID: to.Ptr("DHP_G04_graphPermissionData"),
	// 				Rules: []*armappcomplianceautomation.Rule{
	// 					to.Ptr(armappcomplianceautomation.RuleRequired),
	// 					to.Ptr(armappcomplianceautomation.RuleCharLength)},
	// 					SuperiorQuestionID: to.Ptr("DHP_G07_customerDataProcess"),
	// 				},
	// 				{
	// 					InputType: to.Ptr(armappcomplianceautomation.InputTypeBoolean),
	// 					OptionIDs: []*string{
	// 					},
	// 					QuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 					Rules: []*armappcomplianceautomation.Rule{
	// 						to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 						ShowSubQuestionsValue: to.Ptr("true"),
	// 					},
	// 					{
	// 						InputType: to.Ptr(armappcomplianceautomation.InputTypeText),
	// 						OptionIDs: []*string{
	// 						},
	// 						QuestionID: to.Ptr("DHP_G05_graphPermissionInfo"),
	// 						Rules: []*armappcomplianceautomation.Rule{
	// 							to.Ptr(armappcomplianceautomation.RuleRequired),
	// 							to.Ptr(armappcomplianceautomation.RuleCharLength),
	// 							to.Ptr(armappcomplianceautomation.RulePreventNonEnglishChar)},
	// 							SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 						},
	// 						{
	// 							InputType: to.Ptr(armappcomplianceautomation.InputTypeMultiSelectDropdown),
	// 							OptionIDs: []*string{
	// 								to.Ptr("Croatia"),
	// 								to.Ptr("Cuba"),
	// 								to.Ptr("Curaçao"),
	// 								to.Ptr("Cyprus"),
	// 								to.Ptr("Czechia"),
	// 								to.Ptr("Côte d'Ivoire"),
	// 								to.Ptr("Denmark"),
	// 								to.Ptr("Djibouti"),
	// 								to.Ptr("Dominica"),
	// 								to.Ptr("Dominican Republic (the)"),
	// 								to.Ptr("Ecuador"),
	// 								to.Ptr("Egypt")},
	// 								QuestionID: to.Ptr("DHP_G08_storageLocation"),
	// 								Rules: []*armappcomplianceautomation.Rule{
	// 									to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 									SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 								},
	// 								{
	// 									InputType: to.Ptr(armappcomplianceautomation.InputType("SingleSelectEnum")),
	// 									OptionIDs: []*string{
	// 									},
	// 									QuestionID: to.Ptr("LEG03_complianceDataTermination"),
	// 									Rules: []*armappcomplianceautomation.Rule{
	// 										to.Ptr(armappcomplianceautomation.RuleRequired)},
	// 										SuperiorQuestionID: to.Ptr("DHP_G06_customerDataStorage"),
	// 								}},
	// 							}
}
