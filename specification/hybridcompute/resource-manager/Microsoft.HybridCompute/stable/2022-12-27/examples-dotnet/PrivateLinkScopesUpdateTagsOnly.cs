using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-12-27/examples/PrivateLinkScopesUpdateTagsOnly.json
// this example is just showing the usage of "PrivateLinkScopes_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputePrivateLinkScopeResource created on azure
// for more information of creating HybridComputePrivateLinkScopeResource, please refer to the document of HybridComputePrivateLinkScopeResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string scopeName = "my-privatelinkscope";
ResourceIdentifier hybridComputePrivateLinkScopeResourceId = HybridComputePrivateLinkScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName);
HybridComputePrivateLinkScopeResource hybridComputePrivateLinkScope = client.GetHybridComputePrivateLinkScopeResource(hybridComputePrivateLinkScopeResourceId);

// invoke the operation
HybridComputePrivateLinkScopePatch patch = new HybridComputePrivateLinkScopePatch()
{
    Tags =
    {
    ["Tag1"] = "Value1",
    ["Tag2"] = "Value2",
    },
};
HybridComputePrivateLinkScopeResource result = await hybridComputePrivateLinkScope.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputePrivateLinkScopeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
