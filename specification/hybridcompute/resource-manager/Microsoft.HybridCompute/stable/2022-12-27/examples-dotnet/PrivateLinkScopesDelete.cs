using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-12-27/examples/PrivateLinkScopesDelete.json
// this example is just showing the usage of "PrivateLinkScopes_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputePrivateLinkScopeResource created on azure
// for more information of creating HybridComputePrivateLinkScopeResource, please refer to the document of HybridComputePrivateLinkScopeResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
string resourceGroupName = "my-resource-group";
string scopeName = "my-privatelinkscope";
ResourceIdentifier hybridComputePrivateLinkScopeResourceId = HybridComputePrivateLinkScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName);
HybridComputePrivateLinkScopeResource hybridComputePrivateLinkScope = client.GetHybridComputePrivateLinkScopeResource(hybridComputePrivateLinkScopeResourceId);

// invoke the operation
await hybridComputePrivateLinkScope.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
