package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_UpdateAseNetworkingConfiguration.json
func ExampleEnvironmentsClient_UpdateAseNetworkingConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentsClient().UpdateAseNetworkingConfiguration(ctx, "test-rg", "test-ase", armappservice.AseV3NetworkingConfiguration{
		Properties: &armappservice.AseV3NetworkingConfigurationProperties{
			FtpEnabled:         to.Ptr(true),
			RemoteDebugEnabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AseV3NetworkingConfiguration = armappservice.AseV3NetworkingConfiguration{
	// 	Name: to.Ptr("networking"),
	// 	Type: to.Ptr("Microsoft.Web/hostingEnvironments/configurations/networking"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/configurations/networking"),
	// 	Properties: &armappservice.AseV3NetworkingConfigurationProperties{
	// 		AllowNewPrivateEndpointConnections: to.Ptr(false),
	// 		ExternalInboundIPAddresses: []*string{
	// 			to.Ptr("52.153.248.36")},
	// 			FtpEnabled: to.Ptr(true),
	// 			InternalInboundIPAddresses: []*string{
	// 			},
	// 			LinuxOutboundIPAddresses: []*string{
	// 				to.Ptr("20.88.241.56"),
	// 				to.Ptr("20.88.241.9")},
	// 				RemoteDebugEnabled: to.Ptr(true),
	// 				WindowsOutboundIPAddresses: []*string{
	// 					to.Ptr("20.88.241.56"),
	// 					to.Ptr("20.88.241.9")},
	// 				},
	// 			}
}
