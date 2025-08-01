using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ListInstancePoolOperationsEmpty.json
// this example is just showing the usage of "InstancePoolOperations_ListByInstancePool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InstancePoolResource created on azure
// for more information of creating InstancePoolResource, please refer to the document of InstancePoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resource-group";
string instancePoolName = "test-instance-pool";
ResourceIdentifier instancePoolResourceId = InstancePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instancePoolName);
InstancePoolResource instancePool = client.GetInstancePoolResource(instancePoolResourceId);

// get the collection of this SqlInstancePoolOperationResource
SqlInstancePoolOperationCollection collection = instancePool.GetSqlInstancePoolOperations();

// invoke the operation and iterate over the result
await foreach (SqlInstancePoolOperationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlInstancePoolOperationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
