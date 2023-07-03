using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StorageCache;
using Azure.ResourceManager.StorageCache.Models;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/amlFilesystems_Get.json
// this example is just showing the usage of "amlFilesystems_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AmlFileSystemResource created on azure
// for more information of creating AmlFileSystemResource, please refer to the document of AmlFileSystemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string amlFileSystemName = "fs1";
ResourceIdentifier amlFileSystemResourceId = AmlFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, amlFileSystemName);
AmlFileSystemResource amlFileSystem = client.GetAmlFileSystemResource(amlFileSystemResourceId);

// invoke the operation
AmlFileSystemResource result = await amlFileSystem.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AmlFileSystemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
