using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/SqlAssessmentV2Operations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "SqlAssessmentV2Operations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentGroupResource created on azure
// for more information of creating MigrationAssessmentGroupResource, please refer to the document of MigrationAssessmentGroupResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "rgmigrate";
string projectName = "fci-test6904project";
string groupName = "test_fci_hadr";
ResourceIdentifier migrationAssessmentGroupResourceId = MigrationAssessmentGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName);
MigrationAssessmentGroupResource migrationAssessmentGroup = client.GetMigrationAssessmentGroupResource(migrationAssessmentGroupResourceId);

// get the collection of this MigrationSqlAssessmentV2Resource
MigrationSqlAssessmentV2Collection collection = migrationAssessmentGroup.GetMigrationSqlAssessmentV2s();

// invoke the operation
string assessmentName = "test_swagger_1";
MigrationSqlAssessmentV2Data data = new MigrationSqlAssessmentV2Data
{
    OSLicense = MigrationAssessmentOSLicense.Unknown,
    EnvironmentType = AssessmentEnvironmentType.Production,
    EntityUptime = new AssessmentEntityUptime
    {
        DaysPerMonth = 30,
        HoursPerDay = 24,
    },
    OptimizationLogic = SqlOptimizationLogic.MinimizeCost,
    ReservedInstanceForVm = AssessmentReservedInstance.None,
    AzureOfferCodeForVm = AssessmentOfferCode.MSAZR0003P,
    AzureSqlManagedInstanceSettings = new AssessmentSqlMISettings
    {
        AzureSqlServiceTier = AssessmentSqlServiceTier.Automatic,
        AzureSqlInstanceType = AssessmentSqlInstanceType.SingleInstance,
    },
    AzureSqlDatabaseSettings = new AssessmentSqlDBSettings
    {
        AzureSqlServiceTier = AssessmentSqlServiceTier.Automatic,
        AzureSqlDataBaseType = AssessmentSqlDataBaseType.SingleDatabase,
        AzureSqlComputeTier = MigrationAssessmentComputeTier.Automatic,
        AzureSqlPurchaseModel = AssessmentSqlPurchaseModel.VCore,
    },
    AzureSqlVmInstanceSeries = { AssessmentVmFamily.Eadsv5Series },
    MultiSubnetIntent = MultiSubnetIntent.DisasterRecovery,
    AsyncCommitModeIntent = AsyncCommitModeIntent.DisasterRecovery,
    DisasterRecoveryLocation = new AzureLocation("EastAsia"),
    IsHadrAssessmentEnabled = true,
    ReservedInstance = AssessmentReservedInstance.None,
    SqlServerLicense = AssessmentSqlServerLicense.Unknown,
    AzureLocation = new AzureLocation("SoutheastAsia"),
    AzureOfferCode = AssessmentOfferCode.MSAZR0003P,
    Currency = AssessmentCurrency.USD,
    ScalingFactor = 1,
    Percentile = PercentileOfUtilization.Percentile95,
    TimeRange = AssessmentTimeRange.Day,
    DiscountPercentage = 0,
    SizingCriterion = AssessmentSizingCriterion.PerformanceBased,
};
ArmOperation<MigrationSqlAssessmentV2Resource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, assessmentName, data);
MigrationSqlAssessmentV2Resource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationSqlAssessmentV2Data resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
