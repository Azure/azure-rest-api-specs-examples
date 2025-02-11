using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/ArchiveVersionCreate.json
// this example is just showing the usage of "ArchiveVersions_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryArchiveResource created on azure
// for more information of creating ContainerRegistryArchiveResource, please refer to the document of ContainerRegistryArchiveResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string packageType = "rpm";
string archiveName = "myArchiveName";
ResourceIdentifier containerRegistryArchiveResourceId = ContainerRegistryArchiveResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, packageType, archiveName);
ContainerRegistryArchiveResource containerRegistryArchive = client.GetContainerRegistryArchiveResource(containerRegistryArchiveResourceId);

// get the collection of this ContainerRegistryArchiveVersionResource
ContainerRegistryArchiveVersionCollection collection = containerRegistryArchive.GetContainerRegistryArchiveVersions();

// invoke the operation
string archiveVersionName = "myArchiveVersionName";
ArmOperation<ContainerRegistryArchiveVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, archiveVersionName);
ContainerRegistryArchiveVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryArchiveVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
