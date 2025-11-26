package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetInboundNetworkDependenciesEndpoints.json
func ExampleEnvironmentsClient_NewGetInboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewGetInboundNetworkDependenciesEndpointsPager("Sample-WestUSResourceGroup", "SampleAse", nil)
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
		// page.InboundEnvironmentEndpointCollection = armappservice.InboundEnvironmentEndpointCollection{
		// 	Value: []*armappservice.InboundEnvironmentEndpoint{
		// 		{
		// 			Description: to.Ptr("App Service management"),
		// 			Endpoints: []*string{
		// 				to.Ptr("70.37.57.58/32"),
		// 				to.Ptr("157.55.208.185/32"),
		// 				to.Ptr("23.102.188.65/32"),
		// 				to.Ptr("191.236.154.88/32"),
		// 				to.Ptr("52.174.22.21/32"),
		// 				to.Ptr("13.94.149.179/32"),
		// 				to.Ptr("13.94.143.126/32"),
		// 				to.Ptr("13.94.141.115/32"),
		// 				to.Ptr("52.178.195.197/32"),
		// 				to.Ptr("52.178.190.65/32"),
		// 				to.Ptr("52.178.184.149/32"),
		// 				to.Ptr("52.178.177.147/32"),
		// 				to.Ptr("13.75.127.117/32"),
		// 				to.Ptr("40.83.125.161/32"),
		// 				to.Ptr("40.83.121.56/32"),
		// 				to.Ptr("40.83.120.64/32"),
		// 				to.Ptr("52.187.56.50/32"),
		// 				to.Ptr("52.187.63.37/32"),
		// 				to.Ptr("52.187.59.251/32"),
		// 				to.Ptr("52.187.63.19/32"),
		// 				to.Ptr("52.165.158.140/32"),
		// 				to.Ptr("52.165.152.214/32"),
		// 				to.Ptr("52.165.154.193/32"),
		// 				to.Ptr("52.165.153.122/32"),
		// 				to.Ptr("104.44.129.255/32"),
		// 				to.Ptr("104.44.134.255/32"),
		// 				to.Ptr("104.44.129.243/32"),
		// 				to.Ptr("104.44.129.141/32"),
		// 				to.Ptr("65.52.193.203/32"),
		// 				to.Ptr("70.37.89.222/32"),
		// 				to.Ptr("13.64.115.203/32"),
		// 				to.Ptr("52.225.177.153/32"),
		// 				to.Ptr("65.52.172.237/32")},
		// 				Ports: []*string{
		// 					to.Ptr("454"),
		// 					to.Ptr("455")},
		// 				},
		// 				{
		// 					Description: to.Ptr("App Service Environment VIP"),
		// 					Endpoints: []*string{
		// 						to.Ptr("52.247.209.18/32")},
		// 						Ports: []*string{
		// 							to.Ptr("454"),
		// 							to.Ptr("455"),
		// 							to.Ptr("16001")},
		// 						},
		// 						{
		// 							Description: to.Ptr("App Service Environment subnet"),
		// 							Endpoints: []*string{
		// 								to.Ptr("192.168.250.0/24")},
		// 								Ports: []*string{
		// 									to.Ptr("All")},
		// 							}},
		// 						}
	}
}
