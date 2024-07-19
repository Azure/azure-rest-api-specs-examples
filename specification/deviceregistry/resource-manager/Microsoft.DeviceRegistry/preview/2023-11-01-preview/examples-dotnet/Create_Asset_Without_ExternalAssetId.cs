using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Create_Asset_Without_ExternalAssetId.json
// this example is just showing the usage of "Assets_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DeviceRegistryAssetResource
DeviceRegistryAssetCollection collection = resourceGroupResource.GetDeviceRegistryAssets();

// invoke the operation
string assetName = "my-asset";
DeviceRegistryAssetData data = new DeviceRegistryAssetData(new AzureLocation("West Europe"), new DeviceRegistryExtendedLocation("CustomLocation", "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
{
    AssetType = "MyAssetType",
    Enabled = true,
    DisplayName = "AssetDisplayName",
    Description = "This is a sample Asset",
    AssetEndpointProfileUri = new Uri("https://www.example.com/myAssetEndpointProfile"),
    Manufacturer = "Contoso",
    ManufacturerUri = new Uri("https://www.contoso.com/manufacturerUri"),
    Model = "ContosoModel",
    ProductCode = "SA34VDG",
    HardwareRevision = "1.0",
    SoftwareRevision = "2.0",
    DocumentationUri = new Uri("https://www.example.com/manual"),
    SerialNumber = "64-103816-519918-8",
    DefaultDataPointsConfiguration = "{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}",
    DefaultEventsConfiguration = "{\"publishingInterval\":10,\"samplingInterval\":15,\"queueSize\":20}",
    DataPoints =
    {
    new DataPoint("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt1")
    {
    CapabilityId = "dtmi:com:example:Thermostat:__temperature;1",
    ObservabilityMode = DataPointsObservabilityMode.Counter,
    DataPointConfiguration = "{\"publishingInterval\":8,\"samplingInterval\":8,\"queueSize\":4}",
    },new DataPoint("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt2")
    {
    CapabilityId = "dtmi:com:example:Thermostat:__pressure;1",
    ObservabilityMode = DataPointsObservabilityMode.None,
    DataPointConfiguration = "{\"publishingInterval\":4,\"samplingInterval\":4,\"queueSize\":7}",
    }
    },
    Events =
    {
    new AssetEvent("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt3")
    {
    CapabilityId = "dtmi:com:example:Thermostat:__temperature;1",
    ObservabilityMode = EventsObservabilityMode.None,
    EventConfiguration = "{\"publishingInterval\":7,\"samplingInterval\":1,\"queueSize\":8}",
    },new AssetEvent("nsu=http://microsoft.com/Opc/OpcPlc/;s=FastUInt4")
    {
    CapabilityId = "dtmi:com:example:Thermostat:__pressure;1",
    ObservabilityMode = EventsObservabilityMode.Log,
    EventConfiguration = "{\"publishingInterval\":7,\"samplingInterval\":8,\"queueSize\":4}",
    }
    },
    Tags =
    {
    ["site"] = "building-1",
    },
};
ArmOperation<DeviceRegistryAssetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, assetName, data);
DeviceRegistryAssetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceRegistryAssetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
