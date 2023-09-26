using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseSecurityAlertCreateMin.json
// this example is just showing the usage of "DatabaseSecurityAlertPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseSecurityAlertPolicyResource created on azure
// for more information of creating SqlDatabaseSecurityAlertPolicyResource, please refer to the document of SqlDatabaseSecurityAlertPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-4799";
string serverName = "securityalert-6440";
string databaseName = "testdb";
SqlSecurityAlertPolicyName securityAlertPolicyName = SqlSecurityAlertPolicyName.Default;
ResourceIdentifier sqlDatabaseSecurityAlertPolicyResourceId = SqlDatabaseSecurityAlertPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, securityAlertPolicyName);
SqlDatabaseSecurityAlertPolicyResource sqlDatabaseSecurityAlertPolicy = client.GetSqlDatabaseSecurityAlertPolicyResource(sqlDatabaseSecurityAlertPolicyResourceId);

// invoke the operation
SqlDatabaseSecurityAlertPolicyData data = new SqlDatabaseSecurityAlertPolicyData()
{
    State = SecurityAlertsPolicyState.Enabled,
};
ArmOperation<SqlDatabaseSecurityAlertPolicyResource> lro = await sqlDatabaseSecurityAlertPolicy.UpdateAsync(WaitUntil.Completed, data);
SqlDatabaseSecurityAlertPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlDatabaseSecurityAlertPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
