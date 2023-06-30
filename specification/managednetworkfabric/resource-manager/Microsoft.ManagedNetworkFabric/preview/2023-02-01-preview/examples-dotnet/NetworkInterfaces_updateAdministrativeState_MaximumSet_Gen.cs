using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_updateAdministrativeState_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkInterfaces_updateAdministrativeState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceResource created on azure
// for more information of creating NetworkInterfaceResource, please refer to the document of NetworkInterfaceResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string networkDeviceName = "networkDeviceName";
string networkInterfaceName = "networkInterfaceName";
ResourceIdentifier networkInterfaceResourceId = NetworkInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName, networkInterfaceName);
NetworkInterfaceResource networkInterface = client.GetNetworkInterfaceResource(networkInterfaceResourceId);

// invoke the operation
UpdateAdministrativeState body = new UpdateAdministrativeState()
{
    State = AdministrativeState.Enable,
};
await networkInterface.UpdateAdministrativeStateAsync(WaitUntil.Completed, body);

Console.WriteLine($"Succeeded");
