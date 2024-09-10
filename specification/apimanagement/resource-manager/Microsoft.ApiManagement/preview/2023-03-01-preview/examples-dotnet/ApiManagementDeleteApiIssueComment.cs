using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementDeleteApiIssueComment.json
// this example is just showing the usage of "ApiIssueComment_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiIssueCommentResource created on azure
// for more information of creating ApiIssueCommentResource, please refer to the document of ApiIssueCommentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d1f7558aa04f15146d9d8a";
string issueId = "57d2ef278aa04f0ad01d6cdc";
string commentId = "599e29ab193c3c0bd0b3e2fb";
ResourceIdentifier apiIssueCommentResourceId = ApiIssueCommentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, issueId, commentId);
ApiIssueCommentResource apiIssueComment = client.GetApiIssueCommentResource(apiIssueCommentResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await apiIssueComment.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
