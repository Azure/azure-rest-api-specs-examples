package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListTemplates.json
func ExampleEmailTemplateClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailTemplateClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.EmailTemplateClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.EmailTemplateCollection = armapimanagement.EmailTemplateCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.EmailTemplateContract{
		// 		{
		// 			Name: to.Ptr("ApplicationApprovedNotificationMessage"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/templates"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/templates/ApplicationApprovedNotificationMessage"),
		// 			Properties: &armapimanagement.EmailTemplateContractProperties{
		// 				Description: to.Ptr("Developers who submitted their application for publication in the application gallery on the developer portal receive this email after their submission is approved."),
		// 				Body: to.Ptr("<!DOCTYPE html >\r\n<html>\r\n  <head />\r\n  <body>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n          We are happy to let you know that your request to publish the $AppName application in the application gallery has been approved. Your application has been published and can be viewed <a href=\"http://$DevPortalUrl/Applications/Details/$AppId\">here</a>.\r\n        </p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n  </body>\r\n</html>"),
		// 				IsDefault: to.Ptr(true),
		// 				Parameters: []*armapimanagement.EmailTemplateParametersContractProperties{
		// 					{
		// 						Name: to.Ptr("AppId"),
		// 						Title: to.Ptr("Application id"),
		// 					},
		// 					{
		// 						Name: to.Ptr("AppName"),
		// 						Title: to.Ptr("Application name"),
		// 					},
		// 					{
		// 						Name: to.Ptr("DevFirstName"),
		// 						Title: to.Ptr("Developer first name"),
		// 					},
		// 					{
		// 						Name: to.Ptr("DevLastName"),
		// 						Title: to.Ptr("Developer last name"),
		// 					},
		// 					{
		// 						Name: to.Ptr("OrganizationName"),
		// 						Title: to.Ptr("Organization name"),
		// 					},
		// 					{
		// 						Name: to.Ptr("DevPortalUrl"),
		// 						Title: to.Ptr("Developer portal URL"),
		// 				}},
		// 				Subject: to.Ptr("Your application $AppName is published in the application gallery"),
		// 				Title: to.Ptr("Application gallery submission approved"),
		// 			},
		// 	}},
		// }
	}
}
