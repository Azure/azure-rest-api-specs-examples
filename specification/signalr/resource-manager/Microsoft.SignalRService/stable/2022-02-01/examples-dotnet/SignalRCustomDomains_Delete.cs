using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_Delete.json
// this example is just showing the usage of "SignalRCustomDomains_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SignalRCustomDomainResource created on azure
// for more information of creating SignalRCustomDomainResource, please refer to the document of SignalRCustomDomainResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string name = "example";
ResourceIdentifier signalRCustomDomainResourceId = SignalRCustomDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, name);
SignalRCustomDomainResource signalRCustomDomain = client.GetSignalRCustomDomainResource(signalRCustomDomainResourceId);

// invoke the operation
await signalRCustomDomain.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
