using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/QueueOperationDelete.json
// this example is just showing the usage of "Queue_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageQueueResource created on azure
// for more information of creating StorageQueueResource, please refer to the document of StorageQueueResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
string queueName = "queue6185";
ResourceIdentifier storageQueueResourceId = StorageQueueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, queueName);
StorageQueueResource storageQueue = client.GetStorageQueueResource(storageQueueResourceId);

// invoke the operation
await storageQueue.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
