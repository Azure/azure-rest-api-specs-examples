using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/SqlCollectorOperations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "SqlCollectorOperations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentSqlCollectorResource created on azure
// for more information of creating MigrationAssessmentSqlCollectorResource, please refer to the document of MigrationAssessmentSqlCollectorResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "rgmigrate";
string projectName = "fci-test6904project";
string collectorName = "fci-test0c1esqlsitecollector";
ResourceIdentifier migrationAssessmentSqlCollectorResourceId = MigrationAssessmentSqlCollectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, collectorName);
MigrationAssessmentSqlCollectorResource migrationAssessmentSqlCollector = client.GetMigrationAssessmentSqlCollectorResource(migrationAssessmentSqlCollectorResourceId);

// invoke the operation
MigrationAssessmentSqlCollectorData data = new MigrationAssessmentSqlCollectorData
{
    AgentProperties = new CollectorAgentPropertiesBase
    {
        Id = "630da710-4d44-41f7-a189-72fe3db5502b-agent",
        Version = null,
        LastHeartbeatOn = default,
        SpnDetails = new CollectorAgentSpnPropertiesBase
        {
            Authority = "https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47",
            ApplicationId = "db9c4c3d-477c-4d5a-817b-318276713565",
            Audience = "db9c4c3d-477c-4d5a-817b-318276713565",
            ObjectId = "e50236ad-ad07-47d4-af71-ed7b52d200d5",
            TenantId = Guid.Parse("72f988bf-86f1-41af-91ab-2d7cd011db47"),
        },
    },
    DiscoverySiteId = "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.OffAzure/MasterSites/fci-ankit-test6065mastersite/SqlSites/fci-ankit-test6065sqlsites",
};
ArmOperation<MigrationAssessmentSqlCollectorResource> lro = await migrationAssessmentSqlCollector.UpdateAsync(WaitUntil.Completed, data);
MigrationAssessmentSqlCollectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentSqlCollectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
