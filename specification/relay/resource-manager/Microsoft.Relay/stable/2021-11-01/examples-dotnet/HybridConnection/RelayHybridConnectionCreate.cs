using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionCreate.json
// this example is just showing the usage of "HybridConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RelayNamespaceResource created on azure
// for more information of creating RelayNamespaceResource, please refer to the document of RelayNamespaceResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-01";
ResourceIdentifier relayNamespaceResourceId = RelayNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
RelayNamespaceResource relayNamespace = client.GetRelayNamespaceResource(relayNamespaceResourceId);

// get the collection of this RelayHybridConnectionResource
RelayHybridConnectionCollection collection = relayNamespace.GetRelayHybridConnections();

// invoke the operation
string hybridConnectionName = "example-Relay-Hybrid-01";
RelayHybridConnectionData data = new RelayHybridConnectionData()
{
    IsClientAuthorizationRequired = true,
};
ArmOperation<RelayHybridConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, hybridConnectionName, data);
RelayHybridConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelayHybridConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
