using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionGet.json
// this example is just showing the usage of "DatabaseRecommendedActions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseAdvisorResource created on azure
// for more information of creating SqlDatabaseAdvisorResource, please refer to the document of SqlDatabaseAdvisorResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workloadinsight-demos";
string serverName = "misosisvr";
string databaseName = "IndexAdvisor_test_3";
string advisorName = "CreateIndex";
ResourceIdentifier sqlDatabaseAdvisorResourceId = SqlDatabaseAdvisorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, advisorName);
SqlDatabaseAdvisorResource sqlDatabaseAdvisor = client.GetSqlDatabaseAdvisorResource(sqlDatabaseAdvisorResourceId);

// get the collection of this RecommendedActionResource
RecommendedActionCollection collection = sqlDatabaseAdvisor.GetRecommendedActions();

// invoke the operation
string recommendedActionName = "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB";
NullableResponse<RecommendedActionResource> response = await collection.GetIfExistsAsync(recommendedActionName);
RecommendedActionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RecommendedActionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
