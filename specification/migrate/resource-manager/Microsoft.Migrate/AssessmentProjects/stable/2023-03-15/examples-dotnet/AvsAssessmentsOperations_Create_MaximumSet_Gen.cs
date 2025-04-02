using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AvsAssessmentsOperations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "AvsAssessmentsOperations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentGroupResource created on azure
// for more information of creating MigrationAssessmentGroupResource, please refer to the document of MigrationAssessmentGroupResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawrg";
string projectName = "app18700project";
string groupName = "kuchatur-test";
ResourceIdentifier migrationAssessmentGroupResourceId = MigrationAssessmentGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName);
MigrationAssessmentGroupResource migrationAssessmentGroup = client.GetMigrationAssessmentGroupResource(migrationAssessmentGroupResourceId);

// get the collection of this MigrationAvsAssessmentResource
MigrationAvsAssessmentCollection collection = migrationAssessmentGroup.GetMigrationAvsAssessments();

// invoke the operation
string assessmentName = "asm2";
MigrationAvsAssessmentData data = new MigrationAvsAssessmentData
{
    ProvisioningState = MigrationAssessmentProvisioningState.Succeeded,
    FailuresToTolerateAndRaidLevel = FttAndRaidLevel.Ftt1Raid1,
    VcpuOversubscription = 4,
    NodeType = AssessmentAvsNodeType.Av36,
    ReservedInstance = AssessmentReservedInstance.RI3Year,
    MemOvercommit = 1,
    DedupeCompression = 1.5,
    IsStretchClusterEnabled = true,
    AzureLocation = new AzureLocation("EastUs"),
    AzureOfferCode = AssessmentOfferCode.MSAZR0003P,
    Currency = AssessmentCurrency.USD,
    ScalingFactor = 1,
    Percentile = PercentileOfUtilization.Percentile95,
    TimeRange = AssessmentTimeRange.Day,
    PerfDataStartOn = DateTimeOffset.Parse("2023-09-25T13:35:56.5671462Z"),
    PerfDataEndOn = DateTimeOffset.Parse("2023-09-26T13:35:56.5671462Z"),
    DiscountPercentage = 0,
    SizingCriterion = AssessmentSizingCriterion.AsOnPremises,
};
ArmOperation<MigrationAvsAssessmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, assessmentName, data);
MigrationAvsAssessmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAvsAssessmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
