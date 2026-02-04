using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationDelete.json
// this example is just showing the usage of "NetworkSecurityPerimeterLoggingConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

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
await networkSecurityPerimeterLoggingConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
