using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/EHNameSpaceGet.json
// this example is just showing the usage of "Namespaces_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "SampleSubscription";
string resourceGroupName = "ResurceGroupSample";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EventHubsNamespaceResource
EventHubsNamespaceCollection collection = resourceGroupResource.GetEventHubsNamespaces();

// invoke the operation
string namespaceName = "NamespaceSample";
bool result = await collection.ExistsAsync(namespaceName);

Console.WriteLine($"Succeeded: {result}");
