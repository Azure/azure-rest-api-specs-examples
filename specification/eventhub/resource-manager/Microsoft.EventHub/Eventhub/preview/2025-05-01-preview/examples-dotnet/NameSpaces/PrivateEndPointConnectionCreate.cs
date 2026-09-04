using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/NameSpaces/PrivateEndPointConnectionCreate.json
// this example is just showing the usage of "PrivateEndpointConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsPrivateEndpointConnectionResource created on azure
// for more information of creating EventHubsPrivateEndpointConnectionResource, please refer to the document of EventHubsPrivateEndpointConnectionResource
string subscriptionId = "subID";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-2924";
string privateEndpointConnectionName = "privateEndpointConnectionName";
ResourceIdentifier eventHubsPrivateEndpointConnectionResourceId = EventHubsPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, privateEndpointConnectionName);
EventHubsPrivateEndpointConnectionResource eventHubsPrivateEndpointConnection = client.GetEventHubsPrivateEndpointConnectionResource(eventHubsPrivateEndpointConnectionResourceId);

// invoke the operation
EventHubsPrivateEndpointConnectionData data = new EventHubsPrivateEndpointConnectionData
{
    PrivateEndpointId = new ResourceIdentifier("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"),
    ConnectionState = new EventHubsPrivateLinkServiceConnectionState
    {
        Status = EventHubsPrivateLinkConnectionStatus.Rejected,
        Description = "testing",
    },
    ProvisioningState = EventHubsPrivateEndpointConnectionProvisioningState.Succeeded,
};
ArmOperation<EventHubsPrivateEndpointConnectionResource> lro = await eventHubsPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
EventHubsPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
