using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-09-01/AutonomousDatabases_Update_MaximumSet_Gen.json
// this example is just showing the usage of "AutonomousDatabase_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutonomousDatabaseResource created on azure
// for more information of creating AutonomousDatabaseResource, please refer to the document of AutonomousDatabaseResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgopenapi";
string autonomousdatabasename = "databasedb1";
ResourceIdentifier autonomousDatabaseResourceId = AutonomousDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, autonomousdatabasename);
AutonomousDatabaseResource autonomousDatabase = client.GetAutonomousDatabaseResource(autonomousDatabaseResourceId);

// invoke the operation
AutonomousDatabasePatch patch = new AutonomousDatabasePatch
{
    Tags =
    {
    ["key9827"] = "bygpoqozrwfyiootncgcqq"
    },
    Properties = new AutonomousDatabaseUpdateProperties
    {
        AdminPassword = "<a-password-goes-here>",
        AutonomousMaintenanceScheduleType = AutonomousMaintenanceScheduleType.Early,
        ComputeCount = 56.1F,
        CpuCoreCount = 45,
        CustomerContacts = { new OracleCustomerContact("dummyemail@microsoft.com") },
        DataStorageSizeInTbs = 133,
        DataStorageSizeInGbs = 175271,
        DisplayName = "lrdrjpyyvufnxdzpwvlkmfukpstrjftdxcejcxtnqhxqbhvtzeiokllnspotsqeggddxkjjtf",
        IsAutoScalingEnabled = true,
        IsAutoScalingForStorageEnabled = true,
        PeerDBId = "qmpfwtvpfvbgmulethqznsyyjlpxmyfqfanrymzqsgraavtmlqqbexpzguyqybngoupbshlzpxv",
        IsLocalDataGuardEnabled = true,
        IsMtlsConnectionRequired = true,
        LicenseModel = OracleLicenseModel.LicenseIncluded,
        ScheduledOperationsList = {new ScheduledOperationsTypeUpdate
        {
        DayOfWeekName = OracleDatabaseDayOfWeekName.Monday,
        AutoStartOn = DateTimeOffset.Parse("lwwvkazgmfremfwhckfb"),
        AutoStopOn = DateTimeOffset.Parse("hjwagzxijpiaogulmnmbuqakpqxhkjvaypjqnvbvtjddc"),
        }},
        DatabaseEdition = OracleDatabaseEditionType.StandardEdition,
        LongTermBackupSchedule = new LongTermBackUpScheduleDetails
        {
            RepeatCadence = RepeatCadenceType.OneTime,
            BackupOn = DateTimeOffset.Parse("2025-08-01T04:32:58.715Z"),
            RetentionPeriodInDays = 188,
            IsDisabled = true,
        },
        LocalAdgAutoFailoverMaxDataLossLimit = 212,
        OpenMode = AutonomousDatabaseModeType.ReadOnly,
        PermissionLevel = AutonomousDatabasePermissionLevelType.Restricted,
        Role = DataGuardRoleType.Primary,
        BackupRetentionPeriodInDays = 12,
        WhitelistedIPs = { "kfierlppwurtqrhfxwgfgrnqtmvraignzwsddwmpdijeveuevuoejfmbjvpnlrmmdflilxcwkkzvdofctsdjfxrrrwctihhnchtrouauesqbmlcqhzwnppnhrtitecenlfyshassvajukbwxudhlwixkvkgsessvshcwmleoqujeemwenhwlsccbcjnnviugzgylsxkssalqoicatcvkahogdvweymhdxboyqwhaxuzlmrdbvgbnnetobkbwygcsflzanwknlybvvzgjzjirpfrksbxwgllgfxwdflcisvxpkjecpgdaxccqkzxofedkrawvhzeabmunpykwd" },
    },
};
ArmOperation<AutonomousDatabaseResource> lro = await autonomousDatabase.UpdateAsync(WaitUntil.Completed, patch);
AutonomousDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutonomousDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
