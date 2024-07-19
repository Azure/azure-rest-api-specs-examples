using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StorageActions.Models;
using Azure.ResourceManager.StorageActions;

// Generated from example definition: specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/storageTasksCrud/DeleteStorageTask.json
// this example is just showing the usage of "StorageTasks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageTaskResource created on azure
// for more information of creating StorageTaskResource, please refer to the document of StorageTaskResource
string subscriptionId = "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
string resourceGroupName = "res4228";
string storageTaskName = "mytask1";
ResourceIdentifier storageTaskResourceId = StorageTaskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageTaskName);
StorageTaskResource storageTask = client.GetStorageTaskResource(storageTaskResourceId);

// invoke the operation
await storageTask.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
