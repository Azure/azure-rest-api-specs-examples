using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherConnectionMonitorCreate.json
// this example is just showing the usage of "ConnectionMonitors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkWatcherResource created on azure
// for more information of creating NetworkWatcherResource, please refer to the document of NetworkWatcherResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkWatcherName = "nw1";
ResourceIdentifier networkWatcherResourceId = NetworkWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkWatcherName);
NetworkWatcherResource networkWatcher = client.GetNetworkWatcherResource(networkWatcherResourceId);

// get the collection of this ConnectionMonitorResource
ConnectionMonitorCollection collection = networkWatcher.GetConnectionMonitors();

// invoke the operation
string connectionMonitorName = "cm1";
ConnectionMonitorCreateOrUpdateContent content = new ConnectionMonitorCreateOrUpdateContent
{
    Location = new AzureLocation("eastus"),
    Endpoints = {new ConnectionMonitorEndpoint("source")
    {
    ResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/ct1"),
    }, new ConnectionMonitorEndpoint("destination")
    {
    Address = "bing.com",
    }},
    TestConfigurations = {new ConnectionMonitorTestConfiguration("tcp", ConnectionMonitorTestConfigurationProtocol.Tcp)
    {
    TestFrequencySec = 60,
    TcpConfiguration = new ConnectionMonitorTcpConfiguration
    {
    Port = 80,
    },
    }},
    TestGroups = { new ConnectionMonitorTestGroup("tg", new string[] { "tcp" }, new string[] { "source" }, new string[] { "destination" }) },
};
ArmOperation<ConnectionMonitorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectionMonitorName, content);
ConnectionMonitorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConnectionMonitorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
