package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/SingleSignOnConfigurations_Get.json
func ExampleSingleSignOnConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSingleSignOnConfigurationsClient().Get(ctx, "myResourceGroup", "myMonitor", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SingleSignOnResource = armdatadog.SingleSignOnResource{
	// 	Name: to.Ptr("default"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Datadog/monitors/myMonitor/singleSignOnConfigurations/default"),
	// 	Properties: &armdatadog.SingleSignOnProperties{
	// 		EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		SingleSignOnState: to.Ptr(armdatadog.SingleSignOnStatesEnable),
	// 		SingleSignOnURL: to.Ptr("https://www.datadoghq.com/IAmSomeHash"),
	// 	},
	// }
}
