using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkInterfaces_UpdateAdministrativeState_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkInterfaces_UpdateAdministrativeState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceInterfaceResource created on azure
// for more information of creating NetworkDeviceInterfaceResource, please refer to the document of NetworkDeviceInterfaceResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkDeviceName = "example-device";
string networkInterfaceName = "example-interface";
ResourceIdentifier networkDeviceInterfaceResourceId = NetworkDeviceInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName, networkInterfaceName);
NetworkDeviceInterfaceResource networkDeviceInterface = client.GetNetworkDeviceInterfaceResource(networkDeviceInterfaceResourceId);

// invoke the operation
UpdateAdministrativeStateContent content = new UpdateAdministrativeStateContent()
{
    State = AdministrativeEnableState.Enable,
    ResourceIds =
    {
    new ResourceIdentifier("")
    },
};
ArmOperation<StateUpdateCommonPostActionResult> lro = await networkDeviceInterface.UpdateAdministrativeStateAsync(WaitUntil.Completed, content);
StateUpdateCommonPostActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
