using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/RoutingIntentPut.json
// this example is just showing the usage of "RoutingIntent_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoutingIntentResource created on azure
// for more information of creating RoutingIntentResource, please refer to the document of RoutingIntentResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string routingIntentName = "Intent1";
ResourceIdentifier routingIntentResourceId = RoutingIntentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, routingIntentName);
RoutingIntentResource routingIntent = client.GetRoutingIntentResource(routingIntentResourceId);

// invoke the operation
RoutingIntentData data = new RoutingIntentData
{
    RoutingPolicies = { new RoutingPolicy("InternetTraffic", new string[] { "Internet" }, "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"), new RoutingPolicy("PrivateTrafficPolicy", new string[] { "PrivateTraffic" }, "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1") },
};
ArmOperation<RoutingIntentResource> lro = await routingIntent.UpdateAsync(WaitUntil.Completed, data);
RoutingIntentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoutingIntentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
