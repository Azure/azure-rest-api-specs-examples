package armiotcentral_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Get.json
func ExampleAppsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotcentral.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppsClient().Get(ctx, "resRg", "myIoTCentralApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.App = armiotcentral.App{
	// 	Name: to.Ptr("myIoTCentralApp"),
	// 	Type: to.Ptr("Microsoft.IoTCentral/IoTApps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.IoTCentral/iotApps/myIoTCentralApp"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Identity: &armiotcentral.SystemAssignedServiceIdentity{
	// 		Type: to.Ptr(armiotcentral.SystemAssignedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("8988ab15-3e7a-4586-8a1c-ed07a73a53e9"),
	// 		TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 	},
	// 	Properties: &armiotcentral.AppProperties{
	// 		ApplicationID: to.Ptr("6ebd8fd3-6e34-419e-908f-9be61ec6f6d6"),
	// 		DisplayName: to.Ptr("My IoT Central App"),
	// 		State: to.Ptr(armiotcentral.AppStateCreated),
	// 		Subdomain: to.Ptr("my-iot-central-app"),
	// 		Template: to.Ptr("iotc-pnp-preview@1.0.0"),
	// 	},
	// 	SKU: &armiotcentral.AppSKUInfo{
	// 		Name: to.Ptr(armiotcentral.AppSKU("F1")),
	// 	},
	// }
}
