using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2022-05-10-preview/examples/PrivateLinkScopePrivateLinkResourceGet.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputePrivateLinkScopeResource created on azure
// for more information of creating HybridComputePrivateLinkScopeResource, please refer to the document of HybridComputePrivateLinkScopeResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myResourceGroup";
string scopeName = "myPrivateLinkScope";
ResourceIdentifier hybridComputePrivateLinkScopeResourceId = HybridComputePrivateLinkScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName);
HybridComputePrivateLinkScopeResource hybridComputePrivateLinkScope = client.GetHybridComputePrivateLinkScopeResource(hybridComputePrivateLinkScopeResourceId);

// get the collection of this HybridComputePrivateLinkResource
HybridComputePrivateLinkResourceCollection collection = hybridComputePrivateLinkScope.GetHybridComputePrivateLinkResources();

// invoke the operation
string groupName = "hybridcompute";
bool result = await collection.ExistsAsync(groupName);

Console.WriteLine($"Succeeded: {result}");
