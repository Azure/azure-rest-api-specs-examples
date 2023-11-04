using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/examples/DatabaseGet.json
// this example is just showing the usage of "Databases_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlServerResource created on azure
// for more information of creating MySqlServerResource, please refer to the document of MySqlServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
ResourceIdentifier mySqlServerResourceId = MySqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlServerResource mySqlServer = client.GetMySqlServerResource(mySqlServerResourceId);

// get the collection of this MySqlDatabaseResource
MySqlDatabaseCollection collection = mySqlServer.GetMySqlDatabases();

// invoke the operation
string databaseName = "db1";
NullableResponse<MySqlDatabaseResource> response = await collection.GetIfExistsAsync(databaseName);
MySqlDatabaseResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MySqlDatabaseData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
