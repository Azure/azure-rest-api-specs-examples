using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiIssueAttachment.json
// this example is just showing the usage of "ApiIssueAttachment_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiIssueAttachmentResource created on azure
// for more information of creating ApiIssueAttachmentResource, please refer to the document of ApiIssueAttachmentResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d2ef278aa04f0888cba3f3";
string issueId = "57d2ef278aa04f0ad01d6cdc";
string attachmentId = "57d2ef278aa04f0888cba3f3";
ResourceIdentifier apiIssueAttachmentResourceId = ApiIssueAttachmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, issueId, attachmentId);
ApiIssueAttachmentResource apiIssueAttachment = client.GetApiIssueAttachmentResource(apiIssueAttachmentResourceId);

// invoke the operation
bool result = await apiIssueAttachment.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
