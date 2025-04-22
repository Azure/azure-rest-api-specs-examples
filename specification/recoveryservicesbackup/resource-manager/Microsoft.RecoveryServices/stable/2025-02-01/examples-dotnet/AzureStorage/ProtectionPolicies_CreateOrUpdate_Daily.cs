using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Daily.json
// this example is just showing the usage of "ProtectionPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectionPolicyResource created on azure
// for more information of creating BackupProtectionPolicyResource, please refer to the document of BackupProtectionPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
string vaultName = "swaggertestvault";
string policyName = "dailyPolicy2";
ResourceIdentifier backupProtectionPolicyResourceId = BackupProtectionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, policyName);
BackupProtectionPolicyResource backupProtectionPolicy = client.GetBackupProtectionPolicyResource(backupProtectionPolicyResourceId);

// invoke the operation
BackupProtectionPolicyData data = new BackupProtectionPolicyData(default)
{
    Properties = new FileShareProtectionPolicy
    {
        WorkLoadType = BackupWorkloadType.AzureFileShare,
        SchedulePolicy = new SimpleSchedulePolicy
        {
            ScheduleRunFrequency = ScheduleRunType.Daily,
            ScheduleRunTimes = { default },
        },
        RetentionPolicy = new LongTermRetentionPolicy
        {
            DailySchedule = new DailyRetentionSchedule
            {
                RetentionTimes = { default },
                RetentionDuration = new RetentionDuration
                {
                    Count = 5,
                    DurationType = RetentionDurationType.Days,
                },
            },
            WeeklySchedule = new WeeklyRetentionSchedule
            {
                DaysOfTheWeek = { BackupDayOfWeek.Sunday },
                RetentionTimes = { default },
                RetentionDuration = new RetentionDuration
                {
                    Count = 12,
                    DurationType = RetentionDurationType.Weeks,
                },
            },
            MonthlySchedule = new MonthlyRetentionSchedule
            {
                RetentionScheduleFormatType = RetentionScheduleFormat.Weekly,
                RetentionScheduleWeekly = new WeeklyRetentionFormat
                {
                    DaysOfTheWeek = { BackupDayOfWeek.Sunday },
                    WeeksOfTheMonth = { BackupWeekOfMonth.First },
                },
                RetentionTimes = { default },
                RetentionDuration = new RetentionDuration
                {
                    Count = 60,
                    DurationType = RetentionDurationType.Months,
                },
            },
            YearlySchedule = new YearlyRetentionSchedule
            {
                RetentionScheduleFormatType = RetentionScheduleFormat.Weekly,
                MonthsOfYear = { BackupMonthOfYear.January },
                RetentionScheduleWeekly = new WeeklyRetentionFormat
                {
                    DaysOfTheWeek = { BackupDayOfWeek.Sunday },
                    WeeksOfTheMonth = { BackupWeekOfMonth.First },
                },
                RetentionTimes = { default },
                RetentionDuration = new RetentionDuration
                {
                    Count = 10,
                    DurationType = RetentionDurationType.Years,
                },
            },
        },
        TimeZone = "UTC",
    },
};
ArmOperation<BackupProtectionPolicyResource> lro = await backupProtectionPolicy.UpdateAsync(WaitUntil.Completed, data);
BackupProtectionPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
