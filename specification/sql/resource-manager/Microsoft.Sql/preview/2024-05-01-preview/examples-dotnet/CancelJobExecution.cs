using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/CancelJobExecution.json
// this example is just showing the usage of "JobExecutions_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerJobExecutionResource created on azure
// for more information of creating SqlServerJobExecutionResource, please refer to the document of SqlServerJobExecutionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string serverName = "server1";
string jobAgentName = "agent1";
string jobName = "job1";
Guid jobExecutionId = Guid.Parse("5A86BF65-43AC-F258-2524-9E92992F97CA");
ResourceIdentifier sqlServerJobExecutionResourceId = SqlServerJobExecutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionId);
SqlServerJobExecutionResource sqlServerJobExecution = client.GetSqlServerJobExecutionResource(sqlServerJobExecutionResourceId);

// invoke the operation
await sqlServerJobExecution.CancelAsync();

Console.WriteLine("Succeeded");
