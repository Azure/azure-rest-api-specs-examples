using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyGet.json
// this example is just showing the usage of "LongTermRetentionPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroup";
string serverName = "testserver";
string databaseName = "testDatabase";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// get the collection of this LongTermRetentionPolicyResource
LongTermRetentionPolicyCollection collection = sqlDatabase.GetLongTermRetentionPolicies();

// invoke the operation
LongTermRetentionPolicyName policyName = LongTermRetentionPolicyName.Default;
bool result = await collection.ExistsAsync(policyName);

Console.WriteLine($"Succeeded: {result}");
