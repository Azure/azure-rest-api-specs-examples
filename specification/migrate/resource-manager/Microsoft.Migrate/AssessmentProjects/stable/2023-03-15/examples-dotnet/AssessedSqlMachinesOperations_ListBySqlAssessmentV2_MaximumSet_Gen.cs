using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedSqlMachinesOperations_ListBySqlAssessmentV2_MaximumSet_Gen.json
// this example is just showing the usage of "AssessedSqlMachinesOperations_ListBySqlAssessmentV2" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationSqlAssessmentV2Resource created on azure
// for more information of creating MigrationSqlAssessmentV2Resource, please refer to the document of MigrationSqlAssessmentV2Resource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "rgmigrate";
string projectName = "fci-test6904project";
string groupName = "test_fci_hadr";
string assessmentName = "test_swagger_1";
ResourceIdentifier migrationSqlAssessmentV2ResourceId = MigrationSqlAssessmentV2Resource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName, assessmentName);
MigrationSqlAssessmentV2Resource migrationSqlAssessmentV2 = client.GetMigrationSqlAssessmentV2Resource(migrationSqlAssessmentV2ResourceId);

// get the collection of this MigrationAssessedSqlMachineResource
MigrationAssessedSqlMachineCollection collection = migrationSqlAssessmentV2.GetMigrationAssessedSqlMachines();

// invoke the operation and iterate over the result
string filter = "(contains(Properties/DisplayName,'SQLHAVM17'))";
int? pageSize = 23;
string continuationToken = null;
int? totalRecordCount = 1;
await foreach (MigrationAssessedSqlMachineResource item in collection.GetAllAsync(filter: filter, pageSize: pageSize, continuationToken: continuationToken, totalRecordCount: totalRecordCount))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MigrationAssessedSqlMachineData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
