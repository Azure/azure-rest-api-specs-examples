using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/PrivateEndpointConnections/PrivateEndpointConnectionsDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this RelayPrivateEndpointConnectionResource created on azure
// for more information of creating RelayPrivateEndpointConnectionResource, please refer to the document of RelayPrivateEndpointConnectionResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "myResourceGroup";
string namespaceName = "example-RelayNamespace-5849";
string privateEndpointConnectionName = "{privateEndpointConnection name}";
ResourceIdentifier relayPrivateEndpointConnectionResourceId = RelayPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
RelayPrivateEndpointConnectionResource relayPrivateEndpointConnection = client.GetRelayPrivateEndpointConnectionResource(relayPrivateEndpointConnectionResourceId);

// invoke the operation
await relayPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
