using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Delete_Asset.json
// this example is just showing the usage of "Assets_Delete" operation, for the dependent resources, they will have to be created separately.

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
await deviceRegistryAsset.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
