using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-11-01/Update_AssetEndpointProfile.json
// this example is just showing the usage of "AssetEndpointProfile_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceRegistryAssetEndpointProfileResource created on azure
// for more information of creating DeviceRegistryAssetEndpointProfileResource, please refer to the document of DeviceRegistryAssetEndpointProfileResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string assetEndpointProfileName = "my-assetendpointprofile";
ResourceIdentifier deviceRegistryAssetEndpointProfileResourceId = DeviceRegistryAssetEndpointProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, assetEndpointProfileName);
DeviceRegistryAssetEndpointProfileResource deviceRegistryAssetEndpointProfile = client.GetDeviceRegistryAssetEndpointProfileResource(deviceRegistryAssetEndpointProfileResourceId);

// invoke the operation
DeviceRegistryAssetEndpointProfilePatch patch = new DeviceRegistryAssetEndpointProfilePatch
{
    Properties = new AssetEndpointProfileUpdateProperties
    {
        TargetAddress = new Uri("https://www.example.com/myTargetAddress"),
    },
};
ArmOperation<DeviceRegistryAssetEndpointProfileResource> lro = await deviceRegistryAssetEndpointProfile.UpdateAsync(WaitUntil.Completed, patch);
DeviceRegistryAssetEndpointProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceRegistryAssetEndpointProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
