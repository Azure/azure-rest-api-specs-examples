using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRSharedPrivateLinkResources_Delete.json
// this example is just showing the usage of "SignalRSharedPrivateLinkResources_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SignalRSharedPrivateLinkResource created on azure
// for more information of creating SignalRSharedPrivateLinkResource, please refer to the document of SignalRSharedPrivateLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string sharedPrivateLinkResourceName = "upstream";
ResourceIdentifier signalRSharedPrivateLinkResourceId = SignalRSharedPrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, sharedPrivateLinkResourceName);
SignalRSharedPrivateLinkResource signalRSharedPrivateLinkResource = client.GetSignalRSharedPrivateLinkResource(signalRSharedPrivateLinkResourceId);

// invoke the operation
await signalRSharedPrivateLinkResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
