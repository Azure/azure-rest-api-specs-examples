using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/Databases/stable/2023-12-30/examples/DatabaseCreate.json
// this example is just showing the usage of "Databases_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServerResource created on azure
// for more information of creating MySqlFlexibleServerResource, please refer to the document of MySqlFlexibleServerResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "TestGroup";
string serverName = "testserver";
ResourceIdentifier mySqlFlexibleServerResourceId = MySqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
MySqlFlexibleServerResource mySqlFlexibleServer = client.GetMySqlFlexibleServerResource(mySqlFlexibleServerResourceId);

// get the collection of this MySqlFlexibleServerDatabaseResource
MySqlFlexibleServerDatabaseCollection collection = mySqlFlexibleServer.GetMySqlFlexibleServerDatabases();

// invoke the operation
string databaseName = "db1";
MySqlFlexibleServerDatabaseData data = new MySqlFlexibleServerDatabaseData()
{
    Charset = "utf8",
    Collation = "utf8_general_ci",
};
ArmOperation<MySqlFlexibleServerDatabaseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, databaseName, data);
MySqlFlexibleServerDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MySqlFlexibleServerDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
