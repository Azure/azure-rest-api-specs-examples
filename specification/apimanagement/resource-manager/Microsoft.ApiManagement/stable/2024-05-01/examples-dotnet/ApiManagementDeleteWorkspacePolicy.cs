using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteWorkspacePolicy.json
// this example is just showing the usage of "WorkspacePolicy_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspacePolicyResource created on azure
// for more information of creating ServiceWorkspacePolicyResource, please refer to the document of ServiceWorkspacePolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier serviceWorkspacePolicyResourceId = ServiceWorkspacePolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, policyId);
ServiceWorkspacePolicyResource serviceWorkspacePolicy = client.GetServiceWorkspacePolicyResource(serviceWorkspacePolicyResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await serviceWorkspacePolicy.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine("Succeeded");
