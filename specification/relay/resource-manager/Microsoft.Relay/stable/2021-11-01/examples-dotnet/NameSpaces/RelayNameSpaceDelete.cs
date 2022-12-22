using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Relay;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceDelete.json
// this example is just showing the usage of "Namespaces_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this RelayNamespaceResource created on azure
// for more information of creating RelayNamespaceResource, please refer to the document of RelayNamespaceResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "SouthCentralUS";
string namespaceName = "example-RelayNamespace-5849";
ResourceIdentifier relayNamespaceResourceId = RelayNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
RelayNamespaceResource relayNamespace = client.GetRelayNamespaceResource(relayNamespaceResourceId);

// invoke the operation
await relayNamespace.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
