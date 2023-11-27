using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiIssueAttachment.json
// this example is just showing the usage of "ApiIssueAttachment_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiIssueResource created on azure
// for more information of creating ApiIssueResource, please refer to the document of ApiIssueResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d2ef278aa04f0888cba3f3";
string issueId = "57d2ef278aa04f0ad01d6cdc";
ResourceIdentifier apiIssueResourceId = ApiIssueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, issueId);
ApiIssueResource apiIssue = client.GetApiIssueResource(apiIssueResourceId);

// get the collection of this ApiIssueAttachmentResource
ApiIssueAttachmentCollection collection = apiIssue.GetApiIssueAttachments();

// invoke the operation
string attachmentId = "57d2ef278aa04f0888cba3f3";
NullableResponse<ApiIssueAttachmentResource> response = await collection.GetIfExistsAsync(attachmentId);
ApiIssueAttachmentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiIssueAttachmentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
