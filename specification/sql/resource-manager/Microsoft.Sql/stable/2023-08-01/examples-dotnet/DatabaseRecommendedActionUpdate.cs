using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/DatabaseRecommendedActionUpdate.json
// this example is just showing the usage of "DatabaseRecommendedActions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RecommendedActionResource created on azure
// for more information of creating RecommendedActionResource, please refer to the document of RecommendedActionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workloadinsight-demos";
string serverName = "misosisvr";
string databaseName = "IndexAdvisor_test_3";
string advisorName = "CreateIndex";
string recommendedActionName = "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB";
ResourceIdentifier recommendedActionResourceId = RecommendedActionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName);
RecommendedActionResource recommendedAction = client.GetRecommendedActionResource(recommendedActionResourceId);

// invoke the operation
RecommendedActionData data = new RecommendedActionData
{
    State = new RecommendedActionStateInfo(RecommendedActionCurrentState.Pending),
};
RecommendedActionResource result = await recommendedAction.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RecommendedActionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
