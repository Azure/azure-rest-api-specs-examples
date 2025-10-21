using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/examples/NetworkManagerCommitPost.json
// this example is just showing the usage of "NetworkManagerCommits_Post" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
NetworkManagerCommit networkManagerCommit = new NetworkManagerCommit(new string[] { "useast" }, NetworkConfigurationDeploymentType.SecurityAdmin)
{
    ConfigurationIds = { "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resoureGroupSample/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/SampleSecurityAdminConfig" },
};
ArmOperation<NetworkManagerCommit> lro = await networkManager.PostNetworkManagerCommitAsync(WaitUntil.Completed, networkManagerCommit);
NetworkManagerCommit result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
