using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceTapConfigurationDelete.json
// this example is just showing the usage of "NetworkInterfaceTapConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceTapConfigurationResource created on azure
// for more information of creating NetworkInterfaceTapConfigurationResource, please refer to the document of NetworkInterfaceTapConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string networkInterfaceName = "mynic";
string tapConfigurationName = "tapconfiguration1";
ResourceIdentifier networkInterfaceTapConfigurationResourceId = NetworkInterfaceTapConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkInterfaceName, tapConfigurationName);
NetworkInterfaceTapConfigurationResource networkInterfaceTapConfiguration = client.GetNetworkInterfaceTapConfigurationResource(networkInterfaceTapConfigurationResourceId);

// invoke the operation
await networkInterfaceTapConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
