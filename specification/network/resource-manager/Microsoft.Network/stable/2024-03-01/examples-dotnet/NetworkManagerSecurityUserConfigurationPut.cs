using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/NetworkManagerSecurityUserConfigurationPut.json
// this example is just showing the usage of "SecurityUserConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerResource created on azure
// for more information of creating NetworkManagerResource, please refer to the document of NetworkManagerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
ResourceIdentifier networkManagerResourceId = NetworkManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName);
NetworkManagerResource networkManager = client.GetNetworkManagerResource(networkManagerResourceId);

// get the collection of this NetworkManagerSecurityUserConfigurationResource
NetworkManagerSecurityUserConfigurationCollection collection = networkManager.GetNetworkManagerSecurityUserConfigurations();

// invoke the operation
string configurationName = "myTestSecurityConfig";
NetworkManagerSecurityUserConfigurationData data = new NetworkManagerSecurityUserConfigurationData()
{
    Description = "A sample policy",
};
ArmOperation<NetworkManagerSecurityUserConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationName, data);
NetworkManagerSecurityUserConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkManagerSecurityUserConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
