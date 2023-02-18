using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.SignalR;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_Get.json
// this example is just showing the usage of "SignalRCustomDomains_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SignalRCustomDomainResource created on azure
// for more information of creating SignalRCustomDomainResource, please refer to the document of SignalRCustomDomainResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
string name = "example";
ResourceIdentifier signalRCustomDomainResourceId = SignalRCustomDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, name);
SignalRCustomDomainResource signalRCustomDomain = client.GetSignalRCustomDomainResource(signalRCustomDomainResourceId);

// invoke the operation
SignalRCustomDomainResource result = await signalRCustomDomain.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SignalRCustomDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
