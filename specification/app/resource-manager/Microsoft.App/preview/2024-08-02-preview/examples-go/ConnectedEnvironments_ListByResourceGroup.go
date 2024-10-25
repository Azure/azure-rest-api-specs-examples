package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ConnectedEnvironments_ListByResourceGroup.json
func ExampleConnectedEnvironmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedEnvironmentsClient().NewListByResourceGroupPager("examplerg", nil)
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
		// page.ConnectedEnvironmentCollection = armappcontainers.ConnectedEnvironmentCollection{
		// 	Value: []*armappcontainers.ConnectedEnvironment{
		// 		{
		// 			Name: to.Ptr("sample1"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/sample1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappcontainers.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample1.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample2"),
		// 			Type: to.Ptr("Microsoft.App/connectedEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/sample2"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappcontainers.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr(armappcontainers.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armappcontainers.ConnectedEnvironmentProperties{
		// 				CustomDomainConfiguration: &armappcontainers.CustomDomainConfiguration{
		// 					CustomDomainVerificationID: to.Ptr("custom domain verification id"),
		// 					DNSSuffix: to.Ptr("www.my-name2.com"),
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-06T04:00:00.000Z"); return t}()),
		// 					SubjectName: to.Ptr("CN=www.my-name2.com"),
		// 					Thumbprint: to.Ptr("CERTIFICATE_THUMBPRINT"),
		// 				},
		// 				DefaultDomain: to.Ptr("sample2.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappcontainers.ConnectedEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}
