using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadClassifierMax.json
// this example is just showing the usage of "WorkloadClassifiers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadGroupResource created on azure
// for more information of creating WorkloadGroupResource, please refer to the document of WorkloadGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
string databaseName = "testdb";
string workloadGroupName = "wlm_workloadgroup";
ResourceIdentifier workloadGroupResourceId = WorkloadGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, workloadGroupName);
WorkloadGroupResource workloadGroup = client.GetWorkloadGroupResource(workloadGroupResourceId);

// get the collection of this WorkloadClassifierResource
WorkloadClassifierCollection collection = workloadGroup.GetWorkloadClassifiers();

// invoke the operation
string workloadClassifierName = "wlm_workloadclassifier";
WorkloadClassifierData data = new WorkloadClassifierData()
{
    MemberName = "dbo",
    Label = "test_label",
    Context = "test_context",
    StartTime = "12:00",
    EndTime = "14:00",
    Importance = "high",
};
ArmOperation<WorkloadClassifierResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workloadClassifierName, data);
WorkloadClassifierResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkloadClassifierData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
