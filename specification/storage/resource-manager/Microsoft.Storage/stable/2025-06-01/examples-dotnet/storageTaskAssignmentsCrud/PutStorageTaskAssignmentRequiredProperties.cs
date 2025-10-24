using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/storageTaskAssignmentsCrud/PutStorageTaskAssignmentRequiredProperties.json
// this example is just showing the usage of "StorageTaskAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
string resourceGroupName = "res4228";
string accountName = "sto4445";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// get the collection of this StorageTaskAssignmentResource
StorageTaskAssignmentCollection collection = storageAccount.GetStorageTaskAssignments();

// invoke the operation
string storageTaskAssignmentName = "myassignment1";
StorageTaskAssignmentData data = new StorageTaskAssignmentData(new StorageTaskAssignmentProperties(new ResourceIdentifier("/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"), true, "My Storage task assignment", new StorageTaskAssignmentExecutionContext(new ExecutionTrigger(TaskExecutionTriggerType.RunOnce, new ExecutionTriggerParameters
{
    StartOn = DateTimeOffset.Parse("2022-11-15T21:52:47.8145095Z"),
})), new StorageTaskAssignmentReport("container1")));
ArmOperation<StorageTaskAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, storageTaskAssignmentName, data);
StorageTaskAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageTaskAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
