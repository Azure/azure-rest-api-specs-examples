using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningGet.json
// this example is just showing the usage of "DatabaseAutomaticTuning_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseAutomaticTuningResource created on azure
// for more information of creating SqlDatabaseAutomaticTuningResource, please refer to the document of SqlDatabaseAutomaticTuningResource
string subscriptionId = "c3aa9078-0000-0000-0000-e36f151182d7";
string resourceGroupName = "default-sql-onebox";
string serverName = "testsvr11";
string databaseName = "db1";
ResourceIdentifier sqlDatabaseAutomaticTuningResourceId = SqlDatabaseAutomaticTuningResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseAutomaticTuningResource sqlDatabaseAutomaticTuning = client.GetSqlDatabaseAutomaticTuningResource(sqlDatabaseAutomaticTuningResourceId);

// invoke the operation
SqlDatabaseAutomaticTuningResource result = await sqlDatabaseAutomaticTuning.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlDatabaseAutomaticTuningData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
