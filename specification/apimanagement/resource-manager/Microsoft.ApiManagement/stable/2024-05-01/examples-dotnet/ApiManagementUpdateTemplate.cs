using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateTemplate.json
// this example is just showing the usage of "EmailTemplate_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementEmailTemplateResource created on azure
// for more information of creating ApiManagementEmailTemplateResource, please refer to the document of ApiManagementEmailTemplateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
TemplateName templateName = TemplateName.NewIssueNotificationMessage;
ResourceIdentifier apiManagementEmailTemplateResourceId = ApiManagementEmailTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, templateName);
ApiManagementEmailTemplateResource apiManagementEmailTemplate = client.GetApiManagementEmailTemplateResource(apiManagementEmailTemplateResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementEmailTemplateCreateOrUpdateContent content = new ApiManagementEmailTemplateCreateOrUpdateContent
{
    Subject = "Your request $IssueName was received",
    Body = "<!DOCTYPE html >\r\n<html>\r\n  <head />\r\n  <body>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n          We are happy to let you know that your request to publish the $AppName application in the gallery has been approved. Your application has been published and can be viewed <a href=\"http://$DevPortalUrl/Applications/Details/$AppId\">here</a>.\r\n        </p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n  </body>\r\n</html>",
};
ApiManagementEmailTemplateResource result = await apiManagementEmailTemplate.UpdateAsync(ifMatch, content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementEmailTemplateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
