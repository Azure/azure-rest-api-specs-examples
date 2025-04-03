using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/VmwareCollectorsOperations_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "VmwareCollectorsOperations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentVMwareCollectorResource created on azure
// for more information of creating MigrationAssessmentVMwareCollectorResource, please refer to the document of MigrationAssessmentVMwareCollectorResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawRG";
string projectName = "app18700project";
string vmWareCollectorName = "Vmware2258collector";
ResourceIdentifier migrationAssessmentVMwareCollectorResourceId = MigrationAssessmentVMwareCollectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, vmWareCollectorName);
MigrationAssessmentVMwareCollectorResource migrationAssessmentVMwareCollector = client.GetMigrationAssessmentVMwareCollectorResource(migrationAssessmentVMwareCollectorResourceId);

// invoke the operation
await migrationAssessmentVMwareCollector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
