using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageDiscovery.Models;
using Azure.ResourceManager.StorageDiscovery;

// Generated from example definition: 2025-06-01-preview/StorageDiscoveryWorkspaces_Delete.json
// this example is just showing the usage of "StorageDiscoveryWorkspace_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageDiscoveryWorkspaceResource created on azure
// for more information of creating StorageDiscoveryWorkspaceResource, please refer to the document of StorageDiscoveryWorkspaceResource
string subscriptionId = "b79cb3ba-745e-5d9a-8903-4a02327a7e09";
string resourceGroupName = "sample-rg";
string storageDiscoveryWorkspaceName = "sampleworkspace";
ResourceIdentifier storageDiscoveryWorkspaceResourceId = StorageDiscoveryWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageDiscoveryWorkspaceName);
StorageDiscoveryWorkspaceResource storageDiscoveryWorkspace = client.GetStorageDiscoveryWorkspaceResource(storageDiscoveryWorkspaceResourceId);

// invoke the operation
await storageDiscoveryWorkspace.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
