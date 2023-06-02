using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Hourly.json
// this example is just showing the usage of "ProtectionPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BackupProtectionPolicyResource
string vaultName = "swaggertestvault";
BackupProtectionPolicyCollection collection = resourceGroupResource.GetBackupProtectionPolicies(vaultName);

// invoke the operation
string policyName = "newPolicy2";
BackupProtectionPolicyData data = new BackupProtectionPolicyData(new AzureLocation("placeholder"))
{
    Properties = new FileShareProtectionPolicy()
    {
        WorkLoadType = BackupWorkloadType.AzureFileShare,
        SchedulePolicy = new SimpleSchedulePolicy()
        {
            ScheduleRunFrequency = ScheduleRunType.Hourly,
            HourlySchedule = new BackupHourlySchedule()
            {
                Interval = 4,
                ScheduleWindowStartOn = DateTimeOffset.Parse("2021-09-29T08:00:00.000Z"),
                ScheduleWindowDuration = 12,
            },
        },
        RetentionPolicy = new LongTermRetentionPolicy()
        {
            DailySchedule = new DailyRetentionSchedule()
            {
                RetentionTimes =
                {
                },
                RetentionDuration = new RetentionDuration()
                {
                    Count = 5,
                    DurationType = RetentionDurationType.Days,
                },
            },
            WeeklySchedule = new WeeklyRetentionSchedule()
            {
                DaysOfTheWeek =
                {
                BackupDayOfWeek.Sunday
                },
                RetentionTimes =
                {
                },
                RetentionDuration = new RetentionDuration()
                {
                    Count = 12,
                    DurationType = RetentionDurationType.Weeks,
                },
            },
            MonthlySchedule = new MonthlyRetentionSchedule()
            {
                RetentionScheduleFormatType = RetentionScheduleFormat.Weekly,
                RetentionScheduleWeekly = new WeeklyRetentionFormat()
                {
                    DaysOfTheWeek =
                    {
                    BackupDayOfWeek.Sunday
                    },
                    WeeksOfTheMonth =
                    {
                    BackupWeekOfMonth.First
                    },
                },
                RetentionTimes =
                {
                },
                RetentionDuration = new RetentionDuration()
                {
                    Count = 60,
                    DurationType = RetentionDurationType.Months,
                },
            },
            YearlySchedule = new YearlyRetentionSchedule()
            {
                RetentionScheduleFormatType = RetentionScheduleFormat.Weekly,
                MonthsOfYear =
                {
                BackupMonthOfYear.January
                },
                RetentionScheduleWeekly = new WeeklyRetentionFormat()
                {
                    DaysOfTheWeek =
                    {
                    BackupDayOfWeek.Sunday
                    },
                    WeeksOfTheMonth =
                    {
                    BackupWeekOfMonth.First
                    },
                },
                RetentionTimes =
                {
                },
                RetentionDuration = new RetentionDuration()
                {
                    Count = 10,
                    DurationType = RetentionDurationType.Years,
                },
            },
        },
        TimeZone = "UTC",
    },
};
ArmOperation<BackupProtectionPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyName, data);
BackupProtectionPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
