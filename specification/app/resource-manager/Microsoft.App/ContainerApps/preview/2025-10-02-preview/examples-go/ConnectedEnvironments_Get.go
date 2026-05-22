package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ConnectedEnvironments_Get.json
func ExampleConnectedEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
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
	// res = armappcontainers.ConnectedEnvironmentsClientGetResponse{
	// 	ConnectedEnvironment: armappcontainers.ConnectedEnvironment{
	// 		Name: to.Ptr("examplekenv"),
	// 		Type: to.Ptr("Microsoft.App/connectedEnvironments"),
	// 		ExtendedLocation: &armappcontainers.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 			Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/examplekenv"),
	// 		Location: to.Ptr("North Central US"),
	// 		Properties: &armappcontainers.ConnectedEnvironmentProperties{
	// 			CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
	// 				CustomDomainVerificationID: to.Ptr("custom domain verification id"),
	// 				DNSSuffix: to.Ptr("www.my-name.com"),
	// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00Z"); return t}()),
	// 				SubjectName: to.Ptr("CN=www.my-name.com"),
	// 				Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
	// 			},
	// 			DefaultDomain: to.Ptr("examplekenv.k4apps.io"),
	// 			ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
	// 			StaticIP: to.Ptr("20.42.33.145"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
