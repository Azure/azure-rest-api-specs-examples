using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkDevices_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceResource created on azure
// for more information of creating NetworkDeviceResource, please refer to the document of NetworkDeviceResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkDeviceName = "example-device";
ResourceIdentifier networkDeviceResourceId = NetworkDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName);
NetworkDeviceResource networkDevice = client.GetNetworkDeviceResource(networkDeviceResourceId);

// invoke the operation
NetworkDevicePatch patch = new NetworkDevicePatch()
{
    Annotation = "annotation",
    HostName = "NFA-Device",
    SerialNumber = "Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX",
    Tags =
    {
    ["keyID"] = "KeyValue",
    },
};
ArmOperation<NetworkDeviceResource> lro = await networkDevice.UpdateAsync(WaitUntil.Completed, patch);
NetworkDeviceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkDeviceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
