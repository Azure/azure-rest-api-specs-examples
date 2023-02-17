using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Relay;
using Azure.ResourceManager.Relay.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceGet.json
// this example is just showing the usage of "Namespaces_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RelayNamespaceResource created on azure
// for more information of creating RelayNamespaceResource, please refer to the document of RelayNamespaceResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "RG-eg";
string namespaceName = "example-RelayRelayNamespace-01";
ResourceIdentifier relayNamespaceResourceId = RelayNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
RelayNamespaceResource relayNamespace = client.GetRelayNamespaceResource(relayNamespaceResourceId);

// invoke the operation
RelayNamespaceResource result = await relayNamespace.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelayNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
