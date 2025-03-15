using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/BandwidthSchedulePut.json
// this example is just showing the usage of "BandwidthSchedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeDeviceResource created on azure
// for more information of creating DataBoxEdgeDeviceResource, please refer to the document of DataBoxEdgeDeviceResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
ResourceIdentifier dataBoxEdgeDeviceResourceId = DataBoxEdgeDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName);
DataBoxEdgeDeviceResource dataBoxEdgeDevice = client.GetDataBoxEdgeDeviceResource(dataBoxEdgeDeviceResourceId);

// get the collection of this BandwidthScheduleResource
BandwidthScheduleCollection collection = dataBoxEdgeDevice.GetBandwidthSchedules();

// invoke the operation
string name = "bandwidth-1";
BandwidthScheduleData data = new BandwidthScheduleData(XmlConvert.ToTimeSpan("0:0:0"), XmlConvert.ToTimeSpan("13:59:0"), 100, new DataBoxEdgeDayOfWeek[] { DataBoxEdgeDayOfWeek.Sunday, DataBoxEdgeDayOfWeek.Monday });
ArmOperation<BandwidthScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
BandwidthScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BandwidthScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
