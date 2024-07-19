using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMax.json
// this example is just showing the usage of "ServerDevOpsAuditSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "devAuditTestRG";
string serverName = "devOpsAuditTestSvr";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this SqlServerDevOpsAuditingSettingResource
SqlServerDevOpsAuditingSettingCollection collection = sqlServer.GetSqlServerDevOpsAuditingSettings();

// invoke the operation
string devOpsAuditingSettingsName = "Default";
SqlServerDevOpsAuditingSettingData data = new SqlServerDevOpsAuditingSettingData()
{
    IsAzureMonitorTargetEnabled = true,
    State = BlobAuditingPolicyState.Enabled,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    StorageAccountSubscriptionId = Guid.Parse("00000000-1234-0000-5678-000000000000"),
};
ArmOperation<SqlServerDevOpsAuditingSettingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, devOpsAuditingSettingsName, data);
SqlServerDevOpsAuditingSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerDevOpsAuditingSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
