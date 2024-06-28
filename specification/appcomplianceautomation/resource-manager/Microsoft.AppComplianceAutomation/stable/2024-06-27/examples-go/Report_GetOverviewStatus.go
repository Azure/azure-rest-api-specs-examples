package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_GetOverviewStatus.json
func ExampleProviderActionsClient_GetOverviewStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderActionsClient().GetOverviewStatus(ctx, armappcomplianceautomation.GetOverviewStatusRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GetOverviewStatusResponse = armappcomplianceautomation.GetOverviewStatusResponse{
	// 	StatusList: []*armappcomplianceautomation.StatusItem{
	// 		{
	// 			StatusName: to.Ptr("Active"),
	// 			StatusValue: to.Ptr("100"),
	// 		},
	// 		{
	// 			StatusName: to.Ptr("Failed"),
	// 			StatusValue: to.Ptr("0"),
	// 		},
	// 		{
	// 			StatusName: to.Ptr("Disabled"),
	// 			StatusValue: to.Ptr("0"),
	// 	}},
	// }
}
