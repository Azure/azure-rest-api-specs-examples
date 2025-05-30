package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateTemplate.json
func ExampleEmailTemplateClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailTemplateClient().Update(ctx, "rg1", "apimService1", armapimanagement.TemplateNameNewIssueNotificationMessage, "*", armapimanagement.EmailTemplateUpdateParameters{
		Properties: &armapimanagement.EmailTemplateUpdateParameterProperties{
			Body:    to.Ptr("<!DOCTYPE html >\r\n<html>\r\n  <head />\r\n  <body>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n          We are happy to let you know that your request to publish the $AppName application in the gallery has been approved. Your application has been published and can be viewed <a href=\"http://$DevPortalUrl/Applications/Details/$AppId\">here</a>.\r\n        </p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n  </body>\r\n</html>"),
			Subject: to.Ptr("Your request $IssueName was received"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmailTemplateContract = armapimanagement.EmailTemplateContract{
	// 	Name: to.Ptr("NewIssueNotificationMessage"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/templates"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/templates/NewIssueNotificationMessage"),
	// 	Properties: &armapimanagement.EmailTemplateContractProperties{
	// 		Description: to.Ptr("This email is sent to developers after they create a new topic on the Issues page of the developer portal."),
	// 		Body: to.Ptr("<!DOCTYPE html >\r\n<html>\r\n  <head />\r\n  <body>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Thank you for contacting us. Our API team will review your issue and get back to you soon.</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n          Click this <a href=\"http://$DevPortalUrl/issues/$IssueId\">link</a> to view or edit your request.\r\n        </p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n  </body>\r\n</html>"),
	// 		IsDefault: to.Ptr(true),
	// 		Parameters: []*armapimanagement.EmailTemplateParametersContractProperties{
	// 			{
	// 				Name: to.Ptr("DevFirstName"),
	// 				Title: to.Ptr("Developer first name"),
	// 			},
	// 			{
	// 				Name: to.Ptr("DevLastName"),
	// 				Title: to.Ptr("Developer last name"),
	// 			},
	// 			{
	// 				Name: to.Ptr("IssueId"),
	// 				Title: to.Ptr("Issue id"),
	// 			},
	// 			{
	// 				Name: to.Ptr("IssueName"),
	// 				Title: to.Ptr("Issue name"),
	// 			},
	// 			{
	// 				Name: to.Ptr("OrganizationName"),
	// 				Title: to.Ptr("Organization name"),
	// 			},
	// 			{
	// 				Name: to.Ptr("DevPortalUrl"),
	// 				Title: to.Ptr("Developer portal URL"),
	// 		}},
	// 		Subject: to.Ptr("Your request $IssueName was received"),
	// 		Title: to.Ptr("New issue received"),
	// 	},
	// }
}
