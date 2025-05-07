using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/DatabaseSecurityAlertGet.json
// this example is just showing the usage of "DatabaseSecurityAlertPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-6852";
string serverName = "securityalert-2080";
string databaseName = "testdb";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// get the collection of this SqlDatabaseSecurityAlertPolicyResource
SqlDatabaseSecurityAlertPolicyCollection collection = sqlDatabase.GetSqlDatabaseSecurityAlertPolicies();

// invoke the operation
SqlSecurityAlertPolicyName securityAlertPolicyName = new SqlSecurityAlertPolicyName("default");
NullableResponse<SqlDatabaseSecurityAlertPolicyResource> response = await collection.GetIfExistsAsync(securityAlertPolicyName);
SqlDatabaseSecurityAlertPolicyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlDatabaseSecurityAlertPolicyData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
