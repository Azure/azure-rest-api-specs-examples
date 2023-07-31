using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMin.json
// this example is just showing the usage of "ServerDevOpsAuditSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerDevOpsAuditingSettingResource created on azure
// for more information of creating SqlServerDevOpsAuditingSettingResource, please refer to the document of SqlServerDevOpsAuditingSettingResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "devAuditTestRG";
string serverName = "devOpsAuditTestSvr";
string devOpsAuditingSettingsName = "Default";
ResourceIdentifier sqlServerDevOpsAuditingSettingResourceId = SqlServerDevOpsAuditingSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, devOpsAuditingSettingsName);
SqlServerDevOpsAuditingSettingResource sqlServerDevOpsAuditingSetting = client.GetSqlServerDevOpsAuditingSettingResource(sqlServerDevOpsAuditingSettingResourceId);

// invoke the operation
SqlServerDevOpsAuditingSettingData data = new SqlServerDevOpsAuditingSettingData()
{
    State = BlobAuditingPolicyState.Enabled,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
};
ArmOperation<SqlServerDevOpsAuditingSettingResource> lro = await sqlServerDevOpsAuditingSetting.UpdateAsync(WaitUntil.Completed, data);
SqlServerDevOpsAuditingSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerDevOpsAuditingSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
