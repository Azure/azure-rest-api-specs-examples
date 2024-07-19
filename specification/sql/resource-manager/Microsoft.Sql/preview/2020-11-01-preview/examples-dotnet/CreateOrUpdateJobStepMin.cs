using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMin.json
// this example is just showing the usage of "JobSteps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerJobStepResource created on azure
// for more information of creating SqlServerJobStepResource, please refer to the document of SqlServerJobStepResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string serverName = "server1";
string jobAgentName = "agent1";
string jobName = "job1";
string stepName = "step1";
ResourceIdentifier sqlServerJobStepResourceId = SqlServerJobStepResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, jobAgentName, jobName, stepName);
SqlServerJobStepResource sqlServerJobStep = client.GetSqlServerJobStepResource(sqlServerJobStepResourceId);

// invoke the operation
SqlServerJobStepData data = new SqlServerJobStepData()
{
    TargetGroup = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0",
    Credential = "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0",
    Action = new JobStepAction("select 1"),
};
ArmOperation<SqlServerJobStepResource> lro = await sqlServerJobStep.UpdateAsync(WaitUntil.Completed, data);
SqlServerJobStepResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerJobStepData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
