using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionTargetsByExecution.json
// this example is just showing the usage of "JobTargetExecutions_ListByJobExecution" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerJobResource created on azure
// for more information of creating SqlServerJobResource, please refer to the document of SqlServerJobResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string serverName = "server1";
string jobAgentName = "agent1";
string jobName = "job1";
ResourceIdentifier sqlServerJobResourceId = SqlServerJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, jobAgentName, jobName);
SqlServerJobResource sqlServerJob = client.GetSqlServerJobResource(sqlServerJobResourceId);

// get the collection of this SqlServerJobExecutionResource
SqlServerJobExecutionCollection collection = sqlServerJob.GetSqlServerJobExecutions();

// invoke the operation and iterate over the result
SqlServerJobExecutionCollectionGetJobTargetExecutionsOptions options = new SqlServerJobExecutionCollectionGetJobTargetExecutionsOptions(jobExecutionId: Guid.Parse("5A86BF65-43AC-F258-2524-9E92992F97CA")) { };
await foreach (SqlServerJobExecutionStepTargetResource item in collection.GetJobTargetExecutionsAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlServerJobExecutionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
