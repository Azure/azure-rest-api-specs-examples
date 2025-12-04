using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/NameSpaces/PrivateEndPointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsPrivateEndpointConnectionResource created on azure
// for more information of creating EventHubsPrivateEndpointConnectionResource, please refer to the document of EventHubsPrivateEndpointConnectionResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-3285";
string privateEndpointConnectionName = "928c44d5-b7c6-423b-b6fa-811e0c27b3e0";
ResourceIdentifier eventHubsPrivateEndpointConnectionResourceId = EventHubsPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
EventHubsPrivateEndpointConnectionResource eventHubsPrivateEndpointConnection = client.GetEventHubsPrivateEndpointConnectionResource(eventHubsPrivateEndpointConnectionResourceId);

// invoke the operation
await eventHubsPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
