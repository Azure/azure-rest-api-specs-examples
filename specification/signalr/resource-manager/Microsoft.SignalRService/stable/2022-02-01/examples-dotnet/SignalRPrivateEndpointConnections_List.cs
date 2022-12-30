using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR.Models;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRPrivateEndpointConnections_List.json
// this example is just showing the usage of "SignalRPrivateEndpointConnections_List" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SignalRResource created on azure
// for more information of creating SignalRResource, please refer to the document of SignalRResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
ResourceIdentifier signalRResourceId = SignalRResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
SignalRResource signalR = client.GetSignalRResource(signalRResourceId);

// get the collection of this SignalRPrivateEndpointConnectionResource
SignalRPrivateEndpointConnectionCollection collection = signalR.GetSignalRPrivateEndpointConnections();

// invoke the operation and iterate over the result
await foreach (SignalRPrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SignalRPrivateEndpointConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
