using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_Create_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkInterfaces_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceResource created on azure
// for more information of creating NetworkDeviceResource, please refer to the document of NetworkDeviceResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string networkDeviceName = "networkDeviceName";
ResourceIdentifier networkDeviceResourceId = NetworkDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName);
NetworkDeviceResource networkDevice = client.GetNetworkDeviceResource(networkDeviceResourceId);

// get the collection of this NetworkInterfaceResource
NetworkInterfaceCollection collection = networkDevice.GetNetworkInterfaces();

// invoke the operation
string networkInterfaceName = "networkInterfaceName";
NetworkInterfaceData data = new NetworkInterfaceData()
{
    Annotation = "null",
};
ArmOperation<NetworkInterfaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkInterfaceName, data);
NetworkInterfaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkInterfaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
