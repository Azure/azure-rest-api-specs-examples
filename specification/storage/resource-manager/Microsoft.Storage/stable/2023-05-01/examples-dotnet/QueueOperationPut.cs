using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/QueueOperationPut.json
// this example is just showing the usage of "Queue_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QueueServiceResource created on azure
// for more information of creating QueueServiceResource, please refer to the document of QueueServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
ResourceIdentifier queueServiceResourceId = QueueServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
QueueServiceResource queueService = client.GetQueueServiceResource(queueServiceResourceId);

// get the collection of this StorageQueueResource
StorageQueueCollection collection = queueService.GetStorageQueues();

// invoke the operation
string queueName = "queue6185";
StorageQueueData data = new StorageQueueData();
ArmOperation<StorageQueueResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, queueName, data);
StorageQueueResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageQueueData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
