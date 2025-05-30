using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerScopeConnectionPut.json
// this example is just showing the usage of "ScopeConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerResource created on azure
// for more information of creating NetworkManagerResource, please refer to the document of NetworkManagerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
ResourceIdentifier networkManagerResourceId = NetworkManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName);
NetworkManagerResource networkManager = client.GetNetworkManagerResource(networkManagerResourceId);

// get the collection of this ScopeConnectionResource
ScopeConnectionCollection collection = networkManager.GetScopeConnections();

// invoke the operation
string scopeConnectionName = "TestScopeConnection";
ScopeConnectionData data = new ScopeConnectionData
{
    TenantId = Guid.Parse("6babcaad-604b-40ac-a9d7-9fd97c0b779f"),
    ResourceId = new ResourceIdentifier("subscriptions/f0dc2b34-dfad-40e4-83e0-2309fed8d00b"),
    Description = "This is a scope connection to a cross tenant subscription.",
};
ArmOperation<ScopeConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scopeConnectionName, data);
ScopeConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScopeConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
