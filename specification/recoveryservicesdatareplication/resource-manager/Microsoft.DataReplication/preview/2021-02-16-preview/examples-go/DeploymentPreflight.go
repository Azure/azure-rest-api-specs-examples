package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/DeploymentPreflight.json
func ExampleAzureSiteRecoveryManagementServiceAPIClient_DeploymentPreflight() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureSiteRecoveryManagementServiceAPIClient().DeploymentPreflight(ctx, "rgrecoveryservicesdatareplication", "kjoiahxljomjcmvabaobumg", &armrecoveryservicesdatareplication.AzureSiteRecoveryManagementServiceAPIClientDeploymentPreflightOptions{Body: &armrecoveryservicesdatareplication.DeploymentPreflightModel{
		Resources: []*armrecoveryservicesdatareplication.DeploymentPreflightResource{
			{
				Name:       to.Ptr("xtgugoflfc"),
				Type:       to.Ptr("nsnaptduolqcxsikrewvgjbxqpt"),
				APIVersion: to.Ptr("otihymhvzblycdoxo"),
				Location:   to.Ptr("cbsgtxkjdzwbyp"),
			}},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeploymentPreflightModel = armrecoveryservicesdatareplication.DeploymentPreflightModel{
	// 	Resources: []*armrecoveryservicesdatareplication.DeploymentPreflightResource{
	// 		{
	// 			Name: to.Ptr("xtgugoflfc"),
	// 			Type: to.Ptr("nsnaptduolqcxsikrewvgjbxqpt"),
	// 			APIVersion: to.Ptr("otihymhvzblycdoxo"),
	// 			Location: to.Ptr("cbsgtxkjdzwbyp"),
	// 	}},
	// }
}
