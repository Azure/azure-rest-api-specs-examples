package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementTenantConfigurationSave.json
func ExampleTenantConfigurationClient_BeginSave() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTenantConfigurationClient().BeginSave(ctx, "rg1", "apimService1", armapimanagement.ConfigurationIDNameConfiguration, armapimanagement.SaveConfigurationParameter{
		Properties: &armapimanagement.SaveConfigurationParameterProperties{
			Branch: to.Ptr("master"),
		},
	}, nil)
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
	// res.OperationResultContract = armapimanagement.OperationResultContract{
	// 	Name: to.Ptr("6074e652093a9d0dac3d733c"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/tenant/operationResults"),
	// 	ID: to.Ptr("6074e652093a9d0dac3d733c"),
	// 	Properties: &armapimanagement.OperationResultContractProperties{
	// 		ActionLog: []*armapimanagement.OperationResultLogItemContract{
	// 		},
	// 		ResultInfo: to.Ptr("The configuration was successfully saved to master as commit c0ae274f6046912107bad734834cbf65918668b6."),
	// 		Started: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-13T00:31:14.94Z"); return t}()),
	// 		Status: to.Ptr(armapimanagement.AsyncOperationStatusSucceeded),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-13T00:31:27.59Z"); return t}()),
	// 	},
	// }
}
