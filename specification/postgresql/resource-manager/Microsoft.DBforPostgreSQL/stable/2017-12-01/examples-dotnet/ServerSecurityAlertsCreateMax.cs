using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PostgreSql;
using Azure.ResourceManager.PostgreSql.Models;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMax.json
// this example is just showing the usage of "ServerSecurityAlertPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlServerSecurityAlertPolicyResource created on azure
// for more information of creating PostgreSqlServerSecurityAlertPolicyResource, please refer to the document of PostgreSqlServerSecurityAlertPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securityalert-4799";
string serverName = "securityalert-6440";
PostgreSqlSecurityAlertPolicyName securityAlertPolicyName = PostgreSqlSecurityAlertPolicyName.Default;
ResourceIdentifier postgreSqlServerSecurityAlertPolicyResourceId = PostgreSqlServerSecurityAlertPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, securityAlertPolicyName);
PostgreSqlServerSecurityAlertPolicyResource postgreSqlServerSecurityAlertPolicy = client.GetPostgreSqlServerSecurityAlertPolicyResource(postgreSqlServerSecurityAlertPolicyResourceId);

// invoke the operation
PostgreSqlServerSecurityAlertPolicyData data = new PostgreSqlServerSecurityAlertPolicyData()
{
    State = PostgreSqlServerSecurityAlertPolicyState.Enabled,
    DisabledAlerts =
    {
    "Access_Anomaly","Usage_Anomaly"
    },
    EmailAddresses =
    {
    "testSecurityAlert@microsoft.com"
    },
    SendToEmailAccountAdmins = true,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    RetentionDays = 5,
};
ArmOperation<PostgreSqlServerSecurityAlertPolicyResource> lro = await postgreSqlServerSecurityAlertPolicy.UpdateAsync(WaitUntil.Completed, data);
PostgreSqlServerSecurityAlertPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostgreSqlServerSecurityAlertPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
