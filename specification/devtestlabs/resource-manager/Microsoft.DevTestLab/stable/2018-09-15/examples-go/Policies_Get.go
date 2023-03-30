package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Get.json
func ExamplePoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoliciesClient().Get(ctx, "resourceGroupName", "{labName}", "{policySetName}", "{policyName}", &armdevtestlabs.PoliciesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Policy = armdevtestlabs.Policy{
	// 	Properties: &armdevtestlabs.PolicyProperties{
	// 		Description: to.Ptr("{policyDescription}"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T18:40:48.1739018-07:00"); return t}()),
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
