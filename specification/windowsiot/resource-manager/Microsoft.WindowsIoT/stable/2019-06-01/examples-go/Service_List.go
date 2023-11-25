package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/Service_List.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwindowsiot.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager(nil)
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
		// page.DeviceServiceDescriptionListResult = armwindowsiot.DeviceServiceDescriptionListResult{
		// 	Value: []*armwindowsiot.DeviceService{
		// 		{
		// 			Name: to.Ptr("service1125"),
		// 			Type: to.Ptr("Microsoft.WindowsIoT/Services"),
		// 			ID: to.Ptr("/subscriptions/27de630f-e1ee-42de-8849-90def4986454/resourceGroups/res2627/providers/Microsoft.WindowsIoT/Services/service1125"),
		// 			Properties: &armwindowsiot.DeviceServiceProperties{
		// 				AdminDomainName: to.Ptr("d.e.f"),
		// 				BillingDomainName: to.Ptr("a.b.c"),
		// 				Notes: to.Ptr("blah"),
		// 				Quantity: to.Ptr[int64](1000000),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service3699"),
		// 			Type: to.Ptr("Microsoft.WindowsIoT/Services"),
		// 			ID: to.Ptr("/subscriptions/27de630f-e1ee-42de-8849-90def4986454/resourceGroups/testcmk3/providers/Microsoft.WindowsIoT/Services/service3699"),
		// 			Properties: &armwindowsiot.DeviceServiceProperties{
		// 				AdminDomainName: to.Ptr("d.e.f"),
		// 				BillingDomainName: to.Ptr("a.b.c"),
		// 				Notes: to.Ptr("blah"),
		// 				Quantity: to.Ptr[int64](1000000),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service6637"),
		// 			Type: to.Ptr("Microsoft.WindowsIoT/Services"),
		// 			ID: to.Ptr("/subscriptions/27de630f-e1ee-42de-8849-90def4986454/resourceGroups/testcmk3/providers/Microsoft.WindowsIoT/Services/service6637"),
		// 			Properties: &armwindowsiot.DeviceServiceProperties{
		// 				AdminDomainName: to.Ptr("d.e.f"),
		// 				BillingDomainName: to.Ptr("a.b.c"),
		// 				Notes: to.Ptr("blah"),
		// 				Quantity: to.Ptr[int64](1000000),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service834"),
		// 			Type: to.Ptr("Microsoft.WindowsIoT/Services"),
		// 			ID: to.Ptr("/subscriptions/27de630f-e1ee-42de-8849-90def4986454/resourceGroups/res8186/providers/Microsoft.WindowsIoT/Services/service834"),
		// 			Properties: &armwindowsiot.DeviceServiceProperties{
		// 				AdminDomainName: to.Ptr("d.e.f"),
		// 				BillingDomainName: to.Ptr("a.b.c"),
		// 				Notes: to.Ptr("blah"),
		// 				Quantity: to.Ptr[int64](1000000),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service9174"),
		// 			Type: to.Ptr("Microsoft.WindowsIoT/Services"),
		// 			ID: to.Ptr("/subscriptions/27de630f-e1ee-42de-8849-90def4986454/resourceGroups/testcmk3/providers/Microsoft.WindowsIoT/Services/service9174"),
		// 			Properties: &armwindowsiot.DeviceServiceProperties{
		// 				AdminDomainName: to.Ptr("d.e.f"),
		// 				BillingDomainName: to.Ptr("a.b.c"),
		// 				Notes: to.Ptr("blah"),
		// 				Quantity: to.Ptr[int64](1000000),
		// 				StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "NaN-NaN-NaNTNaN:NaN:NaN.NaNZ"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
