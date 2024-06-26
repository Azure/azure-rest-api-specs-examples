package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Tags.json
func ExampleTagsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTagsClient().Get(ctx, "providers/Microsoft.CostManagement/billingAccounts/1234", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagsResult = armconsumption.TagsResult{
	// 	Name: to.Ptr("tags1"),
	// 	Type: to.Ptr("Microsoft.Consumption/tags"),
	// 	ETag: to.Ptr("\"1d34d012214157f\""),
	// 	ID: to.Ptr("providers/Microsoft.CostManagement/billingAccounts/{billingaccount-id}/providers/Microsoft.Consumption/tags/tags1"),
	// 	Properties: &armconsumption.TagProperties{
	// 		PreviousLink: to.Ptr("https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingaccount-id}/providers/Microsoft.Consumption/tags/?$expand=properties/tags/value&api-version=2021-10-01&startDate=2020-12-01&endDate=2020-12-31&$top=1000&$skiptoken=AQAAAA%3D%3D"),
	// 		Tags: []*armconsumption.Tag{
	// 			{
	// 				Key: to.Ptr("Department"),
	// 			},
	// 			{
	// 				Key: to.Ptr("CostCenter"),
	// 			},
	// 			{
	// 				Key: to.Ptr("Portal"),
	// 			},
	// 			{
	// 				Key: to.Ptr("OrgName"),
	// 			},
	// 			{
	// 				Key: to.Ptr("Namespace"),
	// 			},
	// 			{
	// 				Key: to.Ptr("resourceType"),
	// 			},
	// 			{
	// 				Key: to.Ptr("Subsystem"),
	// 			},
	// 			{
	// 				Key: to.Ptr("Environment"),
	// 			},
	// 			{
	// 				Key: to.Ptr("clusterName"),
	// 		}},
	// 	},
	// }
}
