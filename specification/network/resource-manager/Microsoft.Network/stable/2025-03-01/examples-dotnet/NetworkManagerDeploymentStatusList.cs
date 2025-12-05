using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerDeploymentStatusList.json
// this example is just showing the usage of "NetworkManagerDeploymentStatus_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerResource created on azure
// for more information of creating NetworkManagerResource, please refer to the document of NetworkManagerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resoureGroupSample";
string networkManagerName = "testNetworkManager";
ResourceIdentifier networkManagerResourceId = NetworkManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName);
NetworkManagerResource networkManager = client.GetNetworkManagerResource(networkManagerResourceId);

// invoke the operation and iterate over the result
NetworkManagerDeploymentStatusContent content = new NetworkManagerDeploymentStatusContent
{
    Regions = { "eastus", "westus" },
    DeploymentTypes = { NetworkConfigurationDeploymentType.Connectivity, new NetworkConfigurationDeploymentType("AdminPolicy") },
    SkipToken = "FakeSkipTokenCode",
};
await foreach (NetworkManagerDeploymentStatus item in networkManager.GetNetworkManagerDeploymentStatusAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
