using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDevicePut.json
// this example is just showing the usage of "Devices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataBoxEdgeDeviceResource
DataBoxEdgeDeviceCollection collection = resourceGroupResource.GetDataBoxEdgeDevices();

// invoke the operation
string deviceName = "testedgedevice";
DataBoxEdgeDeviceData data = new DataBoxEdgeDeviceData(new AzureLocation("WUS"))
{
    Sku = new DataBoxEdgeSku
    {
        Name = DataBoxEdgeSkuName.Edge,
        Tier = DataBoxEdgeSkuTier.Standard,
    },
    Tags = { },
};
ArmOperation<DataBoxEdgeDeviceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deviceName, data);
DataBoxEdgeDeviceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeDeviceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
