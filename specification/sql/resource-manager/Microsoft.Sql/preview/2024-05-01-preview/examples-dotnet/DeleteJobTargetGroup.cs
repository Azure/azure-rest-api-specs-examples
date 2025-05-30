using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/DeleteJobTargetGroup.json
// this example is just showing the usage of "JobTargetGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerJobTargetGroupResource created on azure
// for more information of creating SqlServerJobTargetGroupResource, please refer to the document of SqlServerJobTargetGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string serverName = "server1";
string jobAgentName = "agent1";
string targetGroupName = "targetGroup1";
ResourceIdentifier sqlServerJobTargetGroupResourceId = SqlServerJobTargetGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, jobAgentName, targetGroupName);
SqlServerJobTargetGroupResource sqlServerJobTargetGroup = client.GetSqlServerJobTargetGroupResource(sqlServerJobTargetGroupResourceId);

// invoke the operation
await sqlServerJobTargetGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
