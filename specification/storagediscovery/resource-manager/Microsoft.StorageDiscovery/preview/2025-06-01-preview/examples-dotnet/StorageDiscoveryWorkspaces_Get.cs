using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageDiscovery.Models;
using Azure.ResourceManager.StorageDiscovery;

// Generated from example definition: 2025-06-01-preview/StorageDiscoveryWorkspaces_Get.json
// this example is just showing the usage of "StorageDiscoveryWorkspace_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "b79cb3ba-745e-5d9a-8903-4a02327a7e09";
string resourceGroupName = "sample-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StorageDiscoveryWorkspaceResource
StorageDiscoveryWorkspaceCollection collection = resourceGroupResource.GetStorageDiscoveryWorkspaces();

// invoke the operation
string storageDiscoveryWorkspaceName = "Sample-Storage-Workspace";
NullableResponse<StorageDiscoveryWorkspaceResource> response = await collection.GetIfExistsAsync(storageDiscoveryWorkspaceName);
StorageDiscoveryWorkspaceResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StorageDiscoveryWorkspaceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
