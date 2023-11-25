package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListDeploymentsManagementGroupScope.json
func ExampleRemediationsClient_NewListDeploymentsAtManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRemediationsClient().NewListDeploymentsAtManagementGroupPager("financeMg", "myRemediation", &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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
		// page.RemediationDeploymentsListResult = armpolicyinsights.RemediationDeploymentsListResult{
		// 	Value: []*armpolicyinsights.RemediationDeployment{
		// 		{
		// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:51:09.075Z"); return t}()),
		// 			DeploymentID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.resources/deployments/a088e8fd-8600-4126-8d74-fc7ead0e9ae4"),
		// 			LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:52:09.891Z"); return t}()),
		// 			RemediatedResourceID: to.Ptr("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1"),
		// 			ResourceLocation: to.Ptr("eastus"),
		// 			Status: to.Ptr("Succeeded"),
		// 		},
		// 		{
		// 			CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:51:09.075Z"); return t}()),
		// 			DeploymentID: to.Ptr("/subscriptions/c1164d71-0942-499f-bc2f-b6b7b0bae493/resourcegroups/myResourceGroup2/providers/microsoft.resources/deployments/6b8898c0-18a2-449e-987a-8d4d4f634f56"),
		// 			LastUpdatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-09-13T21:52:09.876Z"); return t}()),
		// 			RemediatedResourceID: to.Ptr("/subscriptions/c1164d71-0942-499f-bc2f-b6b7b0bae493/resourcegroups/myResourceGroup2/providers/microsoft.storage/storageaccounts/stor1"),
		// 			ResourceLocation: to.Ptr("westus"),
		// 			Status: to.Ptr("Succeeded"),
		// 	}},
		// }
	}
}
