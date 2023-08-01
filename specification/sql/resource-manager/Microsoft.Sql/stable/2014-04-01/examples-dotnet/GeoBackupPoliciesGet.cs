using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/GeoBackupPoliciesGet.json
// this example is just showing the usage of "GeoBackupPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-4799";
string serverName = "sqlcrudtest-5961";
string databaseName = "testdw";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// get the collection of this GeoBackupPolicyResource
GeoBackupPolicyCollection collection = sqlDatabase.GetGeoBackupPolicies();

// invoke the operation
GeoBackupPolicyName geoBackupPolicyName = GeoBackupPolicyName.Default;
bool result = await collection.ExistsAsync(geoBackupPolicyName);

Console.WriteLine($"Succeeded: {result}");
