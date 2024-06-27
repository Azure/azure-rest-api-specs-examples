package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/TriggerEvaluation.json
func ExampleProviderActionsClient_BeginTriggerEvaluation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProviderActionsClient().BeginTriggerEvaluation(ctx, armappcomplianceautomation.TriggerEvaluationRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TriggerEvaluationResponse = armappcomplianceautomation.TriggerEvaluationResponse{
	// 	Properties: &armappcomplianceautomation.TriggerEvaluationProperty{
	// 		EvaluationEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T05:10:00.000Z"); return t}()),
	// 		QuickAssessments: []*armappcomplianceautomation.QuickAssessment{
	// 			{
	// 				Description: to.Ptr("Protect your storage accounts from potential threats using virtual network rules as a preferred method instead of IP-based filtering. Disabling IP-based filtering prevents public IPs from accessing your storage accounts."),
	// 				DisplayName: to.Ptr("Storage accounts should restrict network access using virtual network rules"),
	// 				RemediationLink: to.Ptr("Protect your storage accounts from potential threats using virtual network rules as a preferred method instead of IP-based filtering. Disabling IP-based filtering prevents public IPs from accessing your storage accounts."),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storettas3iw2megtcarm"),
	// 				ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusHealthy),
	// 				ResponsibilityID: to.Ptr("/providers/microsoft.authorization/policydefinitions/2a1a9cdf-e04d-429a-8416-3bfb72a1b26f"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T05:00:00.000Z"); return t}()),
	// 			},
	// 			{
	// 				Description: to.Ptr(""),
	// 				DisplayName: to.Ptr("Secure transfer to storage accounts should be enabled"),
	// 				RemediationLink: to.Ptr(""),
	// 				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storettas3iw2megtcarm"),
	// 				ResourceStatus: to.Ptr(armappcomplianceautomation.ResourceStatusUnhealthy),
	// 				ResponsibilityID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/404c3081-a854-4457-ae30-26a93ef643f9"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T05:00:00.000Z"); return t}()),
	// 		}},
	// 		ResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storettas3iw2megtcarm")},
	// 			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T05:00:00.000Z"); return t}()),
	// 		},
	// 	}
}
