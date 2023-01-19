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

// this example assumes you already have this StorageQueueResource created on azure
// for more information of creating StorageQueueResource, please refer to the document of StorageQueueResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
string queueName = "queue6185";
ResourceIdentifier storageQueueResourceId = StorageQueueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, queueName);
StorageQueueResource storageQueue = client.GetStorageQueueResource(storageQueueResourceId);

// invoke the operation
StorageQueueResource result = await storageQueue.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageQueueData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
