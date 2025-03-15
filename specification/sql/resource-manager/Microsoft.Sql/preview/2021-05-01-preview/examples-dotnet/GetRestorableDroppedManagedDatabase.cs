using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/GetRestorableDroppedManagedDatabase.json
// this example is just showing the usage of "RestorableDroppedManagedDatabases_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Test1";
string managedInstanceName = "managedInstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// get the collection of this RestorableDroppedManagedDatabaseResource
RestorableDroppedManagedDatabaseCollection collection = managedInstance.GetRestorableDroppedManagedDatabases();

// invoke the operation
string restorableDroppedDatabaseId = "testdb,131403269876900000";
NullableResponse<RestorableDroppedManagedDatabaseResource> response = await collection.GetIfExistsAsync(restorableDroppedDatabaseId);
RestorableDroppedManagedDatabaseResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RestorableDroppedManagedDatabaseData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
