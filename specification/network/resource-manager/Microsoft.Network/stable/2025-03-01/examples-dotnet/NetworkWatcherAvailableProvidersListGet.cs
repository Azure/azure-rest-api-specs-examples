using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherAvailableProvidersListGet.json
// this example is just showing the usage of "NetworkWatchers_ListAvailableProviders" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
AvailableProvidersListContent content = new AvailableProvidersListContent
{
    AzureLocations = { new AzureLocation("West US") },
    Country = "United States",
    State = "washington",
    City = "seattle",
};
ArmOperation<AvailableProvidersList> lro = await networkWatcher.GetAvailableProvidersAsync(WaitUntil.Completed, content);
AvailableProvidersList result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
