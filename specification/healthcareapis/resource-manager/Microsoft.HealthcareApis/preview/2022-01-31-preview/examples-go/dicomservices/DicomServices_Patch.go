package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2f2b633d940230cbbd5bcf1339a2e1c48674e4a2/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2022-01-31-preview/examples/dicomservices/DicomServices_Patch.json
func ExampleDicomServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDicomServicesClient().BeginUpdate(ctx, "testRG", "blue", "workspace1", armhealthcareapis.DicomServicePatchResource{
		Tags: map[string]*string{
			"tagKey": to.Ptr("tagValue"),
		},
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
	// 	Tags: map[string]*string{
	// 		"tagKey": to.Ptr("tagValue"),
	// 	},
	// 	Properties: &armhealthcareapis.DicomServiceProperties{
	// 		AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
	// 			Audiences: []*string{
	// 				to.Ptr("https://azurehealthcareapis.com/"),
	// 				to.Ptr("https://azurehealthcareapis.com")},
	// 				Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
	// 			ServiceURL: to.Ptr("https://workspace1-blue.dicom.azurehealthcareapis.com"),
	// 		},
	// 	}
}
