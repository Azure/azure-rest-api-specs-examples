using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRSharedPrivateLinkResources_Get.json
// this example is just showing the usage of "SignalRSharedPrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SignalRSharedPrivateLinkResource created on azure
// for more information of creating SignalRSharedPrivateLinkResource, please refer to the document of SignalRSharedPrivateLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string sharedPrivateLinkResourceName = "upstream";
ResourceIdentifier signalRSharedPrivateLinkResourceId = SignalRSharedPrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, sharedPrivateLinkResourceName);
SignalRSharedPrivateLinkResource signalRSharedPrivateLinkResource = client.GetSignalRSharedPrivateLinkResource(signalRSharedPrivateLinkResourceId);

// invoke the operation
SignalRSharedPrivateLinkResource result = await signalRSharedPrivateLinkResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SignalRSharedPrivateLinkResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
