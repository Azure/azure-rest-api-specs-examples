using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PrivateEndpointConnections_Delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridTopicPrivateEndpointConnectionResource created on azure
// for more information of creating EventGridTopicPrivateEndpointConnectionResource, please refer to the document of EventGridTopicPrivateEndpointConnectionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string parentName = "exampletopic1";
string privateEndpointConnectionName = "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B";
ResourceIdentifier eventGridTopicPrivateEndpointConnectionResourceId = EventGridTopicPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentName, privateEndpointConnectionName);
EventGridTopicPrivateEndpointConnectionResource eventGridTopicPrivateEndpointConnection = client.GetEventGridTopicPrivateEndpointConnectionResource(eventGridTopicPrivateEndpointConnectionResourceId);

// invoke the operation
await eventGridTopicPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
