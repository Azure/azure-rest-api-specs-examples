using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-09-01-preview/Delete_DiscoveredAsset.json
// this example is just showing the usage of "DiscoveredAsset_Delete" operation, for the dependent resources, they will have to be created separately.

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
await deviceRegistryDiscoveredAsset.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
