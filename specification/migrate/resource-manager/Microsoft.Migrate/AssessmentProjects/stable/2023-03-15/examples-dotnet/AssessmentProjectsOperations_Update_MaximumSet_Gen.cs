using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentProjectsOperations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "AssessmentProjectsOperations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentProjectResource created on azure
// for more information of creating MigrationAssessmentProjectResource, please refer to the document of MigrationAssessmentProjectResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "sakanwar";
string projectName = "sakanwar1204project";
ResourceIdentifier migrationAssessmentProjectResourceId = MigrationAssessmentProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
MigrationAssessmentProjectResource migrationAssessmentProject = client.GetMigrationAssessmentProjectResource(migrationAssessmentProjectResourceId);

// invoke the operation
MigrationAssessmentProjectPatch patch = new MigrationAssessmentProjectPatch
{
    Tags =
    {
    ["Migrate Project"] = "sakanwar-PE-SEA"
    },
    AssessmentSolutionId = new ResourceIdentifier("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
    ProjectStatus = AssessmentProjectStatus.Active,
    CustomerWorkspaceId = null,
    CustomerWorkspaceLocation = null,
    PublicNetworkAccess = "Disabled",
    CustomerStorageAccountArmId = new ResourceIdentifier("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
    ProvisioningState = MigrationAssessmentProvisioningState.Succeeded,
};
ArmOperation<MigrationAssessmentProjectResource> lro = await migrationAssessmentProject.UpdateAsync(WaitUntil.Completed, patch);
MigrationAssessmentProjectResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentProjectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
