using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ListJobStepsByVersion.json
// this example is just showing the usage of "JobSteps_ListByVersion" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerJobVersionResource created on azure
// for more information of creating SqlServerJobVersionResource, please refer to the document of SqlServerJobVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string serverName = "server1";
string jobAgentName = "agent1";
string jobName = "job1";
int jobVersion = 1;
ResourceIdentifier sqlServerJobVersionResourceId = SqlServerJobVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, jobAgentName, jobName, jobVersion);
SqlServerJobVersionResource sqlServerJobVersion = client.GetSqlServerJobVersionResource(sqlServerJobVersionResourceId);

// get the collection of this SqlServerJobVersionStepResource
SqlServerJobVersionStepCollection collection = sqlServerJobVersion.GetSqlServerJobVersionSteps();

// invoke the operation and iterate over the result
await foreach (SqlServerJobVersionStepResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlServerJobStepData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
