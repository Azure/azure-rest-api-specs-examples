package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_CreateOrUpdate.json
func ExamplePoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoliciesClient().CreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{policySetName}", "{policyName}", armdevtestlabs.Policy{
		Location: to.Ptr("{location}"),
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
		Properties: &armdevtestlabs.PolicyProperties{
			Description:   to.Ptr("{policyDescription}"),
			EvaluatorType: to.Ptr(armdevtestlabs.PolicyEvaluatorType("{policyEvaluatorType}")),
			FactData:      to.Ptr("{policyFactData}"),
			FactName:      to.Ptr(armdevtestlabs.PolicyFactName("{policyFactName}")),
			Status:        to.Ptr(armdevtestlabs.PolicyStatus("{policyStatus}")),
			Threshold:     to.Ptr("{policyThreshold}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Policy = armdevtestlabs.Policy{
	// 	Name: to.Ptr("{labName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/policies"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{policyName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Properties: &armdevtestlabs.PolicyProperties{
	// 		Description: to.Ptr("{policyDescription}"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-02T01:40:48.173Z"); return t}()),
	// 		EvaluatorType: to.Ptr(armdevtestlabs.PolicyEvaluatorType("{policyEvaluatorType}")),
	// 		FactData: to.Ptr("{policyFactData}"),
	// 		FactName: to.Ptr(armdevtestlabs.PolicyFactName("{policyFactName}")),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armdevtestlabs.PolicyStatus("{policyStatus}")),
	// 		Threshold: to.Ptr("{policyThreshold}"),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 	},
	// }
}
