using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PrivateEndpointConnections_Get.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridTopicResource created on azure
// for more information of creating EventGridTopicResource, please refer to the document of EventGridTopicResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string parentName = "exampletopic1";
ResourceIdentifier eventGridTopicResourceId = EventGridTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentName);
EventGridTopicResource eventGridTopic = client.GetEventGridTopicResource(eventGridTopicResourceId);

// get the collection of this EventGridTopicPrivateEndpointConnectionResource
EventGridTopicPrivateEndpointConnectionCollection collection = eventGridTopic.GetEventGridTopicPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B";
NullableResponse<EventGridTopicPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
EventGridTopicPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventGridPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
