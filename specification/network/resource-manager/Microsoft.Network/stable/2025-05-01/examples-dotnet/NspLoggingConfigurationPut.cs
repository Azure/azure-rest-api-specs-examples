using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationPut.json
// this example is just showing the usage of "NetworkSecurityPerimeterLoggingConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterResource created on azure
// for more information of creating NetworkSecurityPerimeterResource, please refer to the document of NetworkSecurityPerimeterResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
ResourceIdentifier networkSecurityPerimeterResourceId = NetworkSecurityPerimeterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName);
NetworkSecurityPerimeterResource networkSecurityPerimeter = client.GetNetworkSecurityPerimeterResource(networkSecurityPerimeterResourceId);

// get the collection of this NetworkSecurityPerimeterLoggingConfigurationResource
NetworkSecurityPerimeterLoggingConfigurationCollection collection = networkSecurityPerimeter.GetNetworkSecurityPerimeterLoggingConfigurations();

// invoke the operation
string loggingConfigurationName = "instance";
NetworkSecurityPerimeterLoggingConfigurationData data = new NetworkSecurityPerimeterLoggingConfigurationData
{
    EnabledLogCategories = { "NspPublicInboundPerimeterRulesDenied", "NspPublicOutboundPerimeterRulesDenied" },
};
ArmOperation<NetworkSecurityPerimeterLoggingConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, loggingConfigurationName, data);
NetworkSecurityPerimeterLoggingConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkSecurityPerimeterLoggingConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
