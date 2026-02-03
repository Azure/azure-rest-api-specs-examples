using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerSecurityAdminConfigurationPut_ManualAggregation.json
// this example is just showing the usage of "SecurityAdminConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerResource created on azure
// for more information of creating NetworkManagerResource, please refer to the document of NetworkManagerResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
ResourceIdentifier networkManagerResourceId = NetworkManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName);
NetworkManagerResource networkManager = client.GetNetworkManagerResource(networkManagerResourceId);

// get the collection of this SecurityAdminConfigurationResource
SecurityAdminConfigurationCollection collection = networkManager.GetSecurityAdminConfigurations();

// invoke the operation
string configurationName = "myTestSecurityConfig";
SecurityAdminConfigurationData data = new SecurityAdminConfigurationData
{
    Description = "A configuration which will update any network groups ip addresses at commit times.",
    NetworkGroupAddressSpaceAggregationOption = AddressSpaceAggregationOption.Manual,
};
ArmOperation<SecurityAdminConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationName, data);
SecurityAdminConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAdminConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
