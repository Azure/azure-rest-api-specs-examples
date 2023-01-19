using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/PrivateEndPointConnectionGet.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "subID";
string resourceGroupName = "SDK-EventHub-4794";
string namespaceName = "sdk-Namespace-5828";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// get the collection of this EventHubsPrivateEndpointConnectionResource
EventHubsPrivateEndpointConnectionCollection collection = eventHubsNamespace.GetEventHubsPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "privateEndpointConnectionName";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
