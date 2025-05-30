using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/privateLinkScope/PrivateLinkScopes_Create.json
// this example is just showing the usage of "PrivateLinkScopes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HybridComputePrivateLinkScopeResource
HybridComputePrivateLinkScopeCollection collection = resourceGroupResource.GetHybridComputePrivateLinkScopes();

// invoke the operation
string scopeName = "my-privatelinkscope";
HybridComputePrivateLinkScopeData data = new HybridComputePrivateLinkScopeData(new AzureLocation("westus"));
ArmOperation<HybridComputePrivateLinkScopeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scopeName, data);
HybridComputePrivateLinkScopeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputePrivateLinkScopeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
