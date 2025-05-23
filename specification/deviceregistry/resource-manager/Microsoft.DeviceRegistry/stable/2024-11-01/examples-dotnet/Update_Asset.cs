using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-11-01/Update_Asset.json
// this example is just showing the usage of "Asset_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceRegistryAssetResource created on azure
// for more information of creating DeviceRegistryAssetResource, please refer to the document of DeviceRegistryAssetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string assetName = "my-asset";
ResourceIdentifier deviceRegistryAssetResourceId = DeviceRegistryAssetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, assetName);
DeviceRegistryAssetResource deviceRegistryAsset = client.GetDeviceRegistryAssetResource(deviceRegistryAssetResourceId);

// invoke the operation
DeviceRegistryAssetPatch patch = new DeviceRegistryAssetPatch
{
    Properties = new AssetUpdateProperties
    {
        IsEnabled = false,
        DisplayName = "NewAssetDisplayName",
    },
};
ArmOperation<DeviceRegistryAssetResource> lro = await deviceRegistryAsset.UpdateAsync(WaitUntil.Completed, patch);
DeviceRegistryAssetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceRegistryAssetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
