package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_List.json
func ExampleDicomServicesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDicomServicesClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page.DicomServiceCollection = armhealthcareapis.DicomServiceCollection{
		// 	Value: []*armhealthcareapis.DicomService{
		// 		{
		// 			Name: to.Ptr("blue"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/blue"),
		// 			Properties: &armhealthcareapis.DicomServiceProperties{
		// 				AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
		// 					Audiences: []*string{
		// 						to.Ptr("https://azurehealthcareapis.com/"),
		// 						to.Ptr("https://azurehealthcareapis.com")},
		// 						Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					},
		// 					ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 					ServiceURL: to.Ptr("https://workspace1-blue.dicom.azurehealthcareapis.com"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("red"),
		// 				Type: to.Ptr("Microsoft.HealthcareApis/workspaces/dicomservices"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/dicomservices/red"),
		// 				Properties: &armhealthcareapis.DicomServiceProperties{
		// 					AuthenticationConfiguration: &armhealthcareapis.DicomServiceAuthenticationConfiguration{
		// 						Audiences: []*string{
		// 							to.Ptr("https://azurehealthcareapis.com/"),
		// 							to.Ptr("https://azurehealthcareapis.com")},
		// 							Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 						},
		// 						ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 						ServiceURL: to.Ptr("https://workspace1-red.dicom.azurehealthcareapis.com"),
		// 					},
		// 			}},
		// 		}
	}
}
