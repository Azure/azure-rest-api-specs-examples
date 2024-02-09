using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/LongTermRetentionPolicyCreateOrUpdate.json
// this example is just showing the usage of "LongTermRetentionPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
LongTermRetentionPolicyData data = new LongTermRetentionPolicyData()
{
    MakeBackupsImmutable = true,
    BackupStorageAccessTier = SqlBackupStorageAccessTier.Hot,
    WeeklyRetention = "P1M",
    MonthlyRetention = "P1Y",
    YearlyRetention = "P5Y",
    WeekOfYear = 5,
};
ArmOperation<LongTermRetentionPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyName, data);
LongTermRetentionPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LongTermRetentionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
