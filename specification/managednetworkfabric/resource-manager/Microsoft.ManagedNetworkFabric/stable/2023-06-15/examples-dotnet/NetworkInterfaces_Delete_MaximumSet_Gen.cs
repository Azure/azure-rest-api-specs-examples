using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkInterfaces_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkInterfaces_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceInterfaceResource created on azure
// for more information of creating NetworkDeviceInterfaceResource, please refer to the document of NetworkDeviceInterfaceResource
string subscriptionId = "94D0FD57-C08B-4CA3-A926-6B76D8B7B956";
string resourceGroupName = "rgNetworkDevices";
string networkDeviceName = "sjzd";
string networkInterfaceName = "emrgu";
ResourceIdentifier networkDeviceInterfaceResourceId = NetworkDeviceInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName, networkInterfaceName);
NetworkDeviceInterfaceResource networkDeviceInterface = client.GetNetworkDeviceInterfaceResource(networkDeviceInterfaceResourceId);

// invoke the operation
await networkDeviceInterface.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
