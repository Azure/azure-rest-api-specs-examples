using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkInterfaceIPConfigurationGet.json
// this example is just showing the usage of "NetworkInterfaceIPConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceIPConfigurationResource created on azure
// for more information of creating NetworkInterfaceIPConfigurationResource, please refer to the document of NetworkInterfaceIPConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string networkInterfaceName = "mynic";
string ipConfigurationName = "ipconfig1";
ResourceIdentifier networkInterfaceIPConfigurationResourceId = NetworkInterfaceIPConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkInterfaceName, ipConfigurationName);
NetworkInterfaceIPConfigurationResource networkInterfaceIPConfiguration = client.GetNetworkInterfaceIPConfigurationResource(networkInterfaceIPConfigurationResourceId);

// invoke the operation
NetworkInterfaceIPConfigurationResource result = await networkInterfaceIPConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkInterfaceIPConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
