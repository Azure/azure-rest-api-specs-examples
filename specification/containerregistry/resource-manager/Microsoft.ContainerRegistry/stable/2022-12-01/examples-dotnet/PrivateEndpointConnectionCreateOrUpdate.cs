using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/PrivateEndpointConnectionCreateOrUpdate.json
// this example is just showing the usage of "PrivateEndpointConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryPrivateEndpointConnectionResource created on azure
// for more information of creating ContainerRegistryPrivateEndpointConnectionResource, please refer to the document of ContainerRegistryPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string privateEndpointConnectionName = "myConnection";
ResourceIdentifier containerRegistryPrivateEndpointConnectionResourceId = ContainerRegistryPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, privateEndpointConnectionName);
ContainerRegistryPrivateEndpointConnectionResource containerRegistryPrivateEndpointConnection = client.GetContainerRegistryPrivateEndpointConnectionResource(containerRegistryPrivateEndpointConnectionResourceId);

// invoke the operation
ContainerRegistryPrivateEndpointConnectionData data = new ContainerRegistryPrivateEndpointConnectionData()
{
    ConnectionState = new ContainerRegistryPrivateLinkServiceConnectionState()
    {
        Status = ContainerRegistryPrivateLinkServiceConnectionStatus.Approved,
        Description = "Auto-Approved",
    },
};
ArmOperation<ContainerRegistryPrivateEndpointConnectionResource> lro = await containerRegistryPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
ContainerRegistryPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
