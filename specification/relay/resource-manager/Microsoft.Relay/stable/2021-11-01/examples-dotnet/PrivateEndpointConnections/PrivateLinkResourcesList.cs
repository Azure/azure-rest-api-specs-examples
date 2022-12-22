using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/PrivateEndpointConnections/PrivateLinkResourcesList.json
// this example is just showing the usage of "PrivateLinkResources_List" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this RelayNamespaceResource created on azure
// for more information of creating RelayNamespaceResource, please refer to the document of RelayNamespaceResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "resourcegroup";
string namespaceName = "example-RelayNamespace-5849";
ResourceIdentifier relayNamespaceResourceId = RelayNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
RelayNamespaceResource relayNamespace = client.GetRelayNamespaceResource(relayNamespaceResourceId);

// get the collection of this RelayPrivateLinkResource
RelayPrivateLinkResourceCollection collection = relayNamespace.GetRelayPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (RelayPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RelayPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
