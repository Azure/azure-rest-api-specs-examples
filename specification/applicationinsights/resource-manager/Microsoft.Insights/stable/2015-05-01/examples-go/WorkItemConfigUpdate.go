package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigUpdate.json
func ExampleWorkItemConfigurationsClient_UpdateItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkItemConfigurationsClient().UpdateItem(ctx, "my-resource-group", "my-component", "Visual Studio Team Services", armapplicationinsights.WorkItemCreateConfiguration{
		ConnectorDataConfiguration: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"Custom\":\"{\\\"/fields/System.WorkItemType\\\":\\\"Bug\\\",\\\"/fields/System.AreaPath\\\":\\\"todeletefirst\\\",\\\"/fields/System.AssignedTo\\\":\\\"\\\"}\"}"),
		ConnectorID:                to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
		ValidateOnly:               to.Ptr(true),
		WorkItemProperties: map[string]*string{
			"0": to.Ptr("[object Object]"),
			"1": to.Ptr("[object Object]"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkItemConfiguration = armapplicationinsights.WorkItemConfiguration{
	// 	ConfigDisplayName: to.Ptr("Visual Studio Team Services"),
	// 	ConfigProperties: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"ConfigurationType\":\"Standard\",\"WorkItemType\":\"Bug\",\"AreaPath\":\"todeletefirst\",\"AssignedTo\":\"\"}"),
	// 	ConnectorID: to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
	// 	ID: to.Ptr("Visual Studio Team Services"),
	// 	IsDefault: to.Ptr(true),
	// }
}
