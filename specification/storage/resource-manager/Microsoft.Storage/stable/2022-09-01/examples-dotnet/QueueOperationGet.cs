using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/QueueOperationGet.json
// this example is just showing the usage of "Queue_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(queueName);

Console.WriteLine($"Succeeded: {result}");
