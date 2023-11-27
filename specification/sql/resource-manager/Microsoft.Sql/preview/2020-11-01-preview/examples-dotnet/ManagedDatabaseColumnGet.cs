using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnGet.json
// this example is just showing the usage of "ManagedDatabaseColumns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseTableResource created on azure
// for more information of creating ManagedDatabaseTableResource, please refer to the document of ManagedDatabaseTableResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string managedInstanceName = "myManagedInstanceName";
string databaseName = "myDatabase";
string schemaName = "dbo";
string tableName = "table1";
ResourceIdentifier managedDatabaseTableResourceId = ManagedDatabaseTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName);
ManagedDatabaseTableResource managedDatabaseTable = client.GetManagedDatabaseTableResource(managedDatabaseTableResourceId);

// get the collection of this ManagedDatabaseColumnResource
ManagedDatabaseColumnCollection collection = managedDatabaseTable.GetManagedDatabaseColumns();

// invoke the operation
string columnName = "column1";
NullableResponse<ManagedDatabaseColumnResource> response = await collection.GetIfExistsAsync(columnName);
ManagedDatabaseColumnResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseColumnData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
