using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/ManagedDatabaseTableGet.json
// this example is just showing the usage of "ManagedDatabaseTables_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseSchemaResource created on azure
// for more information of creating ManagedDatabaseSchemaResource, please refer to the document of ManagedDatabaseSchemaResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string managedInstanceName = "myManagedInstanceName";
string databaseName = "myDatabase";
string schemaName = "dbo";
ResourceIdentifier managedDatabaseSchemaResourceId = ManagedDatabaseSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, schemaName);
ManagedDatabaseSchemaResource managedDatabaseSchema = client.GetManagedDatabaseSchemaResource(managedDatabaseSchemaResourceId);

// get the collection of this ManagedDatabaseTableResource
ManagedDatabaseTableCollection collection = managedDatabaseSchema.GetManagedDatabaseTables();

// invoke the operation
string tableName = "table1";
NullableResponse<ManagedDatabaseTableResource> response = await collection.GetIfExistsAsync(tableName);
ManagedDatabaseTableResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseTableData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
