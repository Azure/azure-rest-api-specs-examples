package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironments_Get.json
func ExampleConnectedEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsClient().Get(ctx, "examplerg", "examplekenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironment = armappcontainers.ConnectedEnvironment{
	// 	Name: to.Ptr("examplekenv"),
	// 	Type: to.Ptr("Microsoft.App/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/kubeEnvironments/examplekenv"),
	// 	Location: to.Ptr("North Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappcontainers.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armappcontainers.ConnectedEnvironmentProperties{
	// 		CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 			CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 			DNSSuffix: to.Ptr("www.my-name.com"),
	// 			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 			SubjectName: to.Ptr("CN=www.my-name.com"),
	// 			Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 		},
	// 		DefaultDomain: to.Ptr("examplekenv.k4apps.io"),
	// 		ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("20.42.33.145"),
	// 	},
	// }
}
