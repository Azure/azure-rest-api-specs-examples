using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/VmwareCollectorsOperations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "VmwareCollectorsOperations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentProjectResource created on azure
// for more information of creating MigrationAssessmentProjectResource, please refer to the document of MigrationAssessmentProjectResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawRG";
string projectName = "app18700project";
ResourceIdentifier migrationAssessmentProjectResourceId = MigrationAssessmentProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
MigrationAssessmentProjectResource migrationAssessmentProject = client.GetMigrationAssessmentProjectResource(migrationAssessmentProjectResourceId);

// get the collection of this MigrationAssessmentVMwareCollectorResource
MigrationAssessmentVMwareCollectorCollection collection = migrationAssessmentProject.GetMigrationAssessmentVMwareCollectors();

// invoke the operation
string vmWareCollectorName = "Vmware2258collector";
MigrationAssessmentVMwareCollectorData data = new MigrationAssessmentVMwareCollectorData
{
    ProvisioningState = MigrationAssessmentProvisioningState.Succeeded,
    AgentProperties = new CollectorAgentPropertiesBase
    {
        Id = "fe243486-3318-41fa-aaba-c48b5df75308",
        Version = "1.0.8.383",
        LastHeartbeatOn = DateTimeOffset.Parse("2022-03-29T12:10:08.9167289Z"),
        SpnDetails = new CollectorAgentSpnPropertiesBase
        {
            Authority = "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
            ApplicationId = "82b3e452-c0e8-4662-8347-58282925ae84",
            Audience = "82b3e452-c0e8-4662-8347-58282925ae84",
            ObjectId = "3fc89111-1405-4938-9214-37aa4739401d",
            TenantId = Guid.Parse("72f988bf-86f1-41af-91ab-2d7cd011db47"),
        },
    },
    DiscoverySiteId = "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawRG/providers/Microsoft.OffAzure/VMwareSites/Vmware2744site",
};
ArmOperation<MigrationAssessmentVMwareCollectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vmWareCollectorName, data);
MigrationAssessmentVMwareCollectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentVMwareCollectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
