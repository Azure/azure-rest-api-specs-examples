package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/Service_GetProperties.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwindowsiot.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "res9407", "service8596", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeviceService = armwindowsiot.DeviceService{
	// 	Name: to.Ptr("service8596"),
	// 	Type: to.Ptr("Microsoft.WindowsIoT/Services"),
	// 	ID: to.Ptr("/subscriptions/45b60d85-fd72-427a-a708-f994d26e593e/resourceGroups/res9407/providers/Microsoft.WindowsIoT/Services/service8596"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armwindowsiot.DeviceServiceProperties{
	// 		AdminDomainName: to.Ptr("d.e.f"),
	// 		BillingDomainName: to.Ptr("a.b.c"),
	// 		Notes: to.Ptr("blah"),
	// 		Quantity: to.Ptr[int64](1000000),
	// 		StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
	// 	},
	// }
}
