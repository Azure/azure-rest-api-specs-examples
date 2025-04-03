using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentProjectsOperations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "AssessmentProjectsOperations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "sakanwar";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MigrationAssessmentProjectResource
MigrationAssessmentProjectCollection collection = resourceGroupResource.GetMigrationAssessmentProjects();

// invoke the operation
string projectName = "sakanwar1204project";
MigrationAssessmentProjectData data = new MigrationAssessmentProjectData(new AzureLocation("southeastasia"))
{
    ProvisioningState = MigrationAssessmentProvisioningState.Succeeded,
    AssessmentSolutionId = new ResourceIdentifier("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
    ProjectStatus = AssessmentProjectStatus.Active,
    CustomerWorkspaceId = null,
    CustomerWorkspaceLocation = null,
    PublicNetworkAccess = "Disabled",
    CustomerStorageAccountArmId = new ResourceIdentifier("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa"),
    Tags =
    {
    ["Migrate Project"] = "sakanwar-PE-SEA"
    },
};
ArmOperation<MigrationAssessmentProjectResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, projectName, data);
MigrationAssessmentProjectResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentProjectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
