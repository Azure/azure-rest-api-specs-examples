using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AvsAssessedMachinesOperations_ListByAvsAssessment_MaximumSet_Gen.json
// this example is just showing the usage of "AvsAssessedMachinesOperations_ListByAvsAssessment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAvsAssessmentResource created on azure
// for more information of creating MigrationAvsAssessmentResource, please refer to the document of MigrationAvsAssessmentResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawrg";
string projectName = "app18700project";
string groupName = "kuchatur-test";
string assessmentName = "asm2";
ResourceIdentifier migrationAvsAssessmentResourceId = MigrationAvsAssessmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName, assessmentName);
MigrationAvsAssessmentResource migrationAvsAssessment = client.GetMigrationAvsAssessmentResource(migrationAvsAssessmentResourceId);

// get the collection of this MigrationAvsAssessedMachineResource
MigrationAvsAssessedMachineCollection collection = migrationAvsAssessment.GetMigrationAvsAssessedMachines();

// invoke the operation and iterate over the result
string filter = "ujmwhhuloficljxcjyc";
int? pageSize = 6;
string continuationToken = "qwrjeiukbcicfrkqlqsfukfc";
int? totalRecordCount = 19;
await foreach (MigrationAvsAssessedMachineResource item in collection.GetAllAsync(filter: filter, pageSize: pageSize, continuationToken: continuationToken, totalRecordCount: totalRecordCount))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MigrationAvsAssessedMachineData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
