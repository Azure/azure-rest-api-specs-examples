using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRSharedPrivateLinkResources_Get.json
// this example is just showing the usage of "SignalRSharedPrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SignalRResource created on azure
// for more information of creating SignalRResource, please refer to the document of SignalRResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
ResourceIdentifier signalRResourceId = SignalRResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
SignalRResource signalR = client.GetSignalRResource(signalRResourceId);

// get the collection of this SignalRSharedPrivateLinkResource
SignalRSharedPrivateLinkResourceCollection collection = signalR.GetSignalRSharedPrivateLinkResources();

// invoke the operation
string sharedPrivateLinkResourceName = "upstream";
bool result = await collection.ExistsAsync(sharedPrivateLinkResourceName);

Console.WriteLine($"Succeeded: {result}");
