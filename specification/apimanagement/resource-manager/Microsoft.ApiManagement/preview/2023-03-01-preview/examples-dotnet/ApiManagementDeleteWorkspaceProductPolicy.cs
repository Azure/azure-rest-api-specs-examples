using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementDeleteWorkspaceProductPolicy.json
// this example is just showing the usage of "WorkspaceProductPolicy_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceProductPolicyResource created on azure
// for more information of creating ServiceWorkspaceProductPolicyResource, please refer to the document of ServiceWorkspaceProductPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string productId = "testproduct";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier serviceWorkspaceProductPolicyResourceId = ServiceWorkspaceProductPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, productId, policyId);
ServiceWorkspaceProductPolicyResource serviceWorkspaceProductPolicy = client.GetServiceWorkspaceProductPolicyResource(serviceWorkspaceProductPolicyResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await serviceWorkspaceProductPolicy.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
