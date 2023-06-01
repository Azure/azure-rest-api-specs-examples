using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkInterfaceIPConfigurationGet.json
// this example is just showing the usage of "NetworkInterfaceIPConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceResource created on azure
// for more information of creating NetworkInterfaceResource, please refer to the document of NetworkInterfaceResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string networkInterfaceName = "mynic";
ResourceIdentifier networkInterfaceResourceId = NetworkInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkInterfaceName);
NetworkInterfaceResource networkInterface = client.GetNetworkInterfaceResource(networkInterfaceResourceId);

// get the collection of this NetworkInterfaceIPConfigurationResource
NetworkInterfaceIPConfigurationCollection collection = networkInterface.GetNetworkInterfaceIPConfigurations();

// invoke the operation
string ipConfigurationName = "ipconfig1";
bool result = await collection.ExistsAsync(ipConfigurationName);

Console.WriteLine($"Succeeded: {result}");
