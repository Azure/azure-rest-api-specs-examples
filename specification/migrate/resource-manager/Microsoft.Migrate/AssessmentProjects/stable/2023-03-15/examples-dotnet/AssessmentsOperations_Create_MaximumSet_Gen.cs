using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Migration.Assessment.Models;
using Azure.ResourceManager.Migration.Assessment;

// Generated from example definition: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentsOperations_Create_MaximumSet_Gen.json
// this example is just showing the usage of "AssessmentsOperations_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationAssessmentResource created on azure
// for more information of creating MigrationAssessmentResource, please refer to the document of MigrationAssessmentResource
string subscriptionId = "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
string resourceGroupName = "ayagrawrg";
string projectName = "app18700project";
string groupName = "kuchatur-test";
string assessmentName = "asm1";
ResourceIdentifier migrationAssessmentResourceId = MigrationAssessmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, groupName, assessmentName);
MigrationAssessmentResource migrationAssessment = client.GetMigrationAssessmentResource(migrationAssessmentResourceId);

// invoke the operation
MigrationAssessmentData data = new MigrationAssessmentData
{
    ProvisioningState = MigrationAssessmentProvisioningState.Succeeded,
    EASubscriptionId = "kwsu",
    AzurePricingTier = AssessmentPricingTier.Standard,
    AzureStorageRedundancy = AssessmentStorageRedundancy.Unknown,
    ReservedInstance = AssessmentReservedInstance.None,
    AzureHybridUseBenefit = AssessmentHybridUseBenefit.Unknown,
    AzureDiskTypes = { AssessmentDiskType.Premium, AssessmentDiskType.StandardSsd },
    AzureVmFamilies = { AssessmentVmFamily.DSeries, AssessmentVmFamily.Lsv2Series, AssessmentVmFamily.MSeries, AssessmentVmFamily.Mdsv2Series, AssessmentVmFamily.Msv2Series, AssessmentVmFamily.Mv2Series },
    VmUptime = new AssessmentVmUptime
    {
        DaysPerMonth = 13,
        HoursPerDay = 26,
    },
    AzureLocation = new AzureLocation("njxbwdtsxzhichsnk"),
    AzureOfferCode = AssessmentOfferCode.Unknown,
    Currency = AssessmentCurrency.Unknown,
    ScalingFactor = 24,
    Percentile = PercentileOfUtilization.Percentile50,
    TimeRange = AssessmentTimeRange.Day,
    PerfDataStartOn = DateTimeOffset.Parse("2023-09-26T09:36:48.491Z"),
    PerfDataEndOn = DateTimeOffset.Parse("2023-09-26T09:36:48.491Z"),
    DiscountPercentage = 6,
    SizingCriterion = AssessmentSizingCriterion.PerformanceBased,
};
ArmOperation<MigrationAssessmentResource> lro = await migrationAssessment.UpdateAsync(WaitUntil.Completed, data);
MigrationAssessmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MigrationAssessmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
