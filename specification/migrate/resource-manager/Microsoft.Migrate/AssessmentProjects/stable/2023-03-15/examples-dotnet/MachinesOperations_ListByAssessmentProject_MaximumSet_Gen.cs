using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/MachinesOperations_ListByAssessmentProject_MaximumSet_Gen.json
// this example is just showing the usage of "MachinesOperations_ListByAssessmentProject" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentProjectResource created on azure
// for more information of creating MigrationAssessmentProjectResource, please refer to the document of MigrationAssessmentProjectResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawrg";
string projectName = "app18700project";
ResourceIdentifier migrationAssessmentProjectResourceId = MigrationAssessmentProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
MigrationAssessmentProjectResource migrationAssessmentProject = client.GetMigrationAssessmentProjectResource(migrationAssessmentProjectResourceId);

// get the collection of this MigrationAssessmentMachineResource
MigrationAssessmentMachineCollection collection = migrationAssessmentProject.GetMigrationAssessmentMachines();

// invoke the operation and iterate over the result
string filter = null;
int? pageSize = 1;
string continuationToken = null;
int? totalRecordCount = 1;
await foreach (MigrationAssessmentMachineResource item in collection.GetAllAsync(filter: filter, pageSize: pageSize, continuationToken: continuationToken, totalRecordCount: totalRecordCount))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MigrationAssessmentMachineData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
