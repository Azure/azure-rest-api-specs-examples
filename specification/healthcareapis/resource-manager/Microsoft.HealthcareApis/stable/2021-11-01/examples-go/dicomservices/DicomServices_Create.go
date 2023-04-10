package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Create.json
func ExampleDicomServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDicomServicesClient().BeginCreateOrUpdate(ctx, "testRG", "workspace1", "blue", armhealthcareapis.DicomService{
		Location:   to.Ptr("westus"),
		Properties: &armhealthcareapis.DicomServiceProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DicomService = armhealthcareapis.DicomService{
	// 	Name: to.Ptr("blue"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/blue"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armhealthcareapis.DicomServiceProperties{
	// 		AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
	// 			Audiences: []*string{
	// 				to.Ptr("https://azurehealthcareapis.com/"),
	// 				to.Ptr("https://azurehealthcareapis.com")},
	// 				Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
