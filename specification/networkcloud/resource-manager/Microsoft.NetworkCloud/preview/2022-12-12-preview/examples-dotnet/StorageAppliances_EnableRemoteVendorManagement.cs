using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/StorageAppliances_EnableRemoteVendorManagement.json
// this example is just showing the usage of "StorageAppliances_EnableRemoteVendorManagement" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageApplianceResource created on azure
// for more information of creating StorageApplianceResource, please refer to the document of StorageApplianceResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string storageApplianceName = "storageApplianceName";
ResourceIdentifier storageApplianceResourceId = StorageApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageApplianceName);
StorageApplianceResource storageAppliance = client.GetStorageApplianceResource(storageApplianceResourceId);

// invoke the operation
StorageApplianceEnableRemoteVendorManagementContent content = new StorageApplianceEnableRemoteVendorManagementContent()
{
    SupportEndpoints =
    {
    "10.0.0.0/24"
    },
};
await storageAppliance.EnableRemoteVendorManagementAsync(WaitUntil.Completed, content: content);

Console.WriteLine($"Succeeded");
