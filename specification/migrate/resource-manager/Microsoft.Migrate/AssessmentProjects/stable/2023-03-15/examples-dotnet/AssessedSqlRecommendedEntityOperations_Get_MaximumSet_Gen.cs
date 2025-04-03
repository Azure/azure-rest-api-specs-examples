using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedSqlRecommendedEntityOperations_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AssessedSqlRecommendedEntityOperations_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this MigrationAssessedSqlRecommendedEntityResource
MigrationAssessedSqlRecommendedEntityCollection collection = migrationSqlAssessmentV2.GetMigrationAssessedSqlRecommendedEntities();

// invoke the operation
string recommendedAssessedEntityName = "cc64c9dc-b38e-435d-85ad-d509df5d92c6";
NullableResponse<MigrationAssessedSqlRecommendedEntityResource> response = await collection.GetIfExistsAsync(recommendedAssessedEntityName);
MigrationAssessedSqlRecommendedEntityResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MigrationAssessedSqlRecommendedEntityData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
