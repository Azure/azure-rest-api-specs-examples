using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Dell.Storage.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Dell.Storage;

// Generated from example definition: 2025-03-21-preview/FileSystems_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "FileSystemResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DellFileSystemResource created on azure
// for more information of creating DellFileSystemResource, please refer to the document of DellFileSystemResource
string subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
string resourceGroupName = "rgDell";
string filesystemName = "abcd";
ResourceIdentifier dellFileSystemResourceId = DellFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, filesystemName);
DellFileSystemResource dellFileSystem = client.GetDellFileSystemResource(dellFileSystemResourceId);

// invoke the operation
await dellFileSystem.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
