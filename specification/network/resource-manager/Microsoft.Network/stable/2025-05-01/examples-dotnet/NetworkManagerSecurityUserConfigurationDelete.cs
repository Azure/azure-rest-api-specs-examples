using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerSecurityUserConfigurationDelete.json
// this example is just showing the usage of "SecurityUserConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerSecurityUserConfigurationResource created on azure
// for more information of creating NetworkManagerSecurityUserConfigurationResource, please refer to the document of NetworkManagerSecurityUserConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string configurationName = "myTestSecurityConfig";
ResourceIdentifier networkManagerSecurityUserConfigurationResourceId = NetworkManagerSecurityUserConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, configurationName);
NetworkManagerSecurityUserConfigurationResource networkManagerSecurityUserConfiguration = client.GetNetworkManagerSecurityUserConfigurationResource(networkManagerSecurityUserConfigurationResourceId);

// invoke the operation
bool? force = false;
await networkManagerSecurityUserConfiguration.DeleteAsync(WaitUntil.Completed, force: force);

Console.WriteLine("Succeeded");
