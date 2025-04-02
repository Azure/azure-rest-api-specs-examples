using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/SqlCollectorOperations_Get_MaximumSet_Gen.json
// this example is just showing the usage of "SqlCollectorOperations_Get" operation, for the dependent resources, they will have to be created separately.

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
MigrationAssessmentSqlCollectorResource result = await migrationAssessmentSqlCollector.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentSqlCollectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
