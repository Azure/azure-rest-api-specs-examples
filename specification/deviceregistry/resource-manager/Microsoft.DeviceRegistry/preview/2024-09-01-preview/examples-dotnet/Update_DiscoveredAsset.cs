using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-09-01-preview/Update_DiscoveredAsset.json
// this example is just showing the usage of "DiscoveredAsset_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceRegistryDiscoveredAssetResource created on azure
// for more information of creating DeviceRegistryDiscoveredAssetResource, please refer to the document of DeviceRegistryDiscoveredAssetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string discoveredAssetName = "my-discoveredasset";
ResourceIdentifier deviceRegistryDiscoveredAssetResourceId = DeviceRegistryDiscoveredAssetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, discoveredAssetName);
DeviceRegistryDiscoveredAssetResource deviceRegistryDiscoveredAsset = client.GetDeviceRegistryDiscoveredAssetResource(deviceRegistryDiscoveredAssetResourceId);

// invoke the operation
DeviceRegistryDiscoveredAssetPatch patch = new DeviceRegistryDiscoveredAssetPatch
{
    Properties = new DiscoveredAssetUpdateProperties
    {
        DocumentationUri = new Uri("https://www.example.com/manual-2"),
        DefaultTopic = new DeviceRegistryTopic("/path/defaultTopic")
        {
            Retain = DeviceRegistryTopicRetainType.Never,
        },
    },
};
ArmOperation<DeviceRegistryDiscoveredAssetResource> lro = await deviceRegistryDiscoveredAsset.UpdateAsync(WaitUntil.Completed, patch);
DeviceRegistryDiscoveredAssetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceRegistryDiscoveredAssetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
