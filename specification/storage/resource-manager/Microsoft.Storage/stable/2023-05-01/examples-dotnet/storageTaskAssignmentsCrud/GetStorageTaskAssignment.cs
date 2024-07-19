using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/storageTaskAssignmentsCrud/GetStorageTaskAssignment.json
// this example is just showing the usage of "StorageTaskAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageTaskAssignmentResource created on azure
// for more information of creating StorageTaskAssignmentResource, please refer to the document of StorageTaskAssignmentResource
string subscriptionId = "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
string resourceGroupName = "res4228";
string accountName = "sto4445";
string storageTaskAssignmentName = "myassignment1";
ResourceIdentifier storageTaskAssignmentResourceId = StorageTaskAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, storageTaskAssignmentName);
StorageTaskAssignmentResource storageTaskAssignment = client.GetStorageTaskAssignmentResource(storageTaskAssignmentResourceId);

// invoke the operation
StorageTaskAssignmentResource result = await storageTaskAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageTaskAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
