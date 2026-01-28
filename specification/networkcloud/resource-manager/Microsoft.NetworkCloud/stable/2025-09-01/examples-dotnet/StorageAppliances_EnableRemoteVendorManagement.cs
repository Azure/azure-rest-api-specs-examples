using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/StorageAppliances_EnableRemoteVendorManagement.json
// this example is just showing the usage of "StorageAppliances_EnableRemoteVendorManagement" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudStorageApplianceResource created on azure
// for more information of creating NetworkCloudStorageApplianceResource, please refer to the document of NetworkCloudStorageApplianceResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string storageApplianceName = "storageApplianceName";
ResourceIdentifier networkCloudStorageApplianceResourceId = NetworkCloudStorageApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageApplianceName);
NetworkCloudStorageApplianceResource networkCloudStorageAppliance = client.GetNetworkCloudStorageApplianceResource(networkCloudStorageApplianceResourceId);

// invoke the operation
StorageApplianceEnableRemoteVendorManagementContent content = new StorageApplianceEnableRemoteVendorManagementContent
{
    SupportEndpoints = { "10.0.0.0/24" },
};
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudStorageAppliance.EnableRemoteVendorManagementAsync(WaitUntil.Completed, content: content);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
