package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_GetVipInfo.json
func ExampleEnvironmentsClient_GetVipInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentsClient().GetVipInfo(ctx, "test-rg", "test-ase", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AddressResponse = armappservice.AddressResponse{
	// 	Name: to.Ptr("test-ase"),
	// 	Type: to.Ptr("Microsoft.Web/hostingEnvironments/capacities"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/Microsoft.Web/hostingEnvironments/test-ase/capacities/virtualip"),
	// 	Properties: &armappservice.AddressResponseProperties{
	// 		OutboundIPAddresses: []*string{
	// 			to.Ptr("20.112.141.120")},
	// 			ServiceIPAddress: to.Ptr("20.112.141.120"),
	// 			VipMappings: []*armappservice.VirtualIPMapping{
	// 				{
	// 					InUse: to.Ptr(false),
	// 					InternalHTTPPort: to.Ptr[int32](20003),
	// 					InternalHTTPSPort: to.Ptr[int32](20001),
	// 					VirtualIP: to.Ptr("20.112.141.135"),
	// 				},
	// 				{
	// 					InUse: to.Ptr(false),
	// 					InternalHTTPPort: to.Ptr[int32](20004),
	// 					InternalHTTPSPort: to.Ptr[int32](20002),
	// 					VirtualIP: to.Ptr("20.112.141.150"),
	// 			}},
	// 		},
	// 	}
}
