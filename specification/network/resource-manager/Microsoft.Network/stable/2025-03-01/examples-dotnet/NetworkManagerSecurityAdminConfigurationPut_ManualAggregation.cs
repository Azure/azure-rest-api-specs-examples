using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerSecurityAdminConfigurationPut_ManualAggregation.json
// this example is just showing the usage of "SecurityAdminConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAdminConfigurationResource created on azure
// for more information of creating SecurityAdminConfigurationResource, please refer to the document of SecurityAdminConfigurationResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
ResourceIdentifier securityAdminConfigurationResourceId = SecurityAdminConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName);
SecurityAdminConfigurationResource securityAdminConfiguration = client.GetSecurityAdminConfigurationResource(securityAdminConfigurationResourceId);

// invoke the operation
SecurityAdminConfigurationData data = new SecurityAdminConfigurationData
{
    Description = "A configuration which will update any network groups ip addresses at commit times.",
    NetworkGroupAddressSpaceAggregationOption = AddressSpaceAggregationOption.Manual,
};
ArmOperation<SecurityAdminConfigurationResource> lro = await securityAdminConfiguration.UpdateAsync(WaitUntil.Completed, data);
SecurityAdminConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAdminConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
