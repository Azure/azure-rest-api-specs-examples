using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLoggingConfigurationPut.json
// this example is just showing the usage of "NetworkSecurityPerimeterLoggingConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterLoggingConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterLoggingConfigurationResource, please refer to the document of NetworkSecurityPerimeterLoggingConfigurationResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
string loggingConfigurationName = "instance";
ResourceIdentifier networkSecurityPerimeterLoggingConfigurationResourceId = NetworkSecurityPerimeterLoggingConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName, loggingConfigurationName);
NetworkSecurityPerimeterLoggingConfigurationResource networkSecurityPerimeterLoggingConfiguration = client.GetNetworkSecurityPerimeterLoggingConfigurationResource(networkSecurityPerimeterLoggingConfigurationResourceId);

// invoke the operation
NetworkSecurityPerimeterLoggingConfigurationData data = new NetworkSecurityPerimeterLoggingConfigurationData
{
    EnabledLogCategories = { "NspPublicInboundPerimeterRulesDenied", "NspPublicOutboundPerimeterRulesDenied" },
};
ArmOperation<NetworkSecurityPerimeterLoggingConfigurationResource> lro = await networkSecurityPerimeterLoggingConfiguration.UpdateAsync(WaitUntil.Completed, data);
NetworkSecurityPerimeterLoggingConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkSecurityPerimeterLoggingConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
