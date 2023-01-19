using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/CreateOrUpdateSqlVirtualMachineAutomatedBackupWeekly.json
// this example is just showing the usage of "SqlVirtualMachines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SqlVmResource
SqlVmCollection collection = resourceGroupResource.GetSqlVms();

// invoke the operation
string sqlVmName = "testvm";
SqlVmData data = new SqlVmData(new AzureLocation("northeurope"))
{
    VirtualMachineResourceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
    SqlServerLicenseType = SqlServerLicenseType.Payg,
    SqlManagement = SqlManagementMode.Full,
    SqlImageSku = SqlImageSku.Enterprise,
    AutoPatchingSettings = new SqlVmAutoPatchingSettings()
    {
        IsEnabled = true,
        DayOfWeek = SqlVmAutoPatchingDayOfWeek.Sunday,
        MaintenanceWindowStartingHour = 2,
        MaintenanceWindowDurationInMinutes = 60,
    },
    AutoBackupSettings = new SqlVmAutoBackupSettings()
    {
        IsEnabled = true,
        IsEncryptionEnabled = true,
        RetentionPeriodInDays = 17,
        StorageAccountUri = new Uri("https://teststorage.blob.core.windows.net/"),
        StorageContainerName = "testcontainer",
        StorageAccessKey = "<primary storage access key>",
        Password = "<Password>",
        AreSystemDbsIncludedInBackup = true,
        BackupScheduleType = SqVmBackupScheduleType.Manual,
        FullBackupFrequency = SqlVmFullBackupFrequency.Weekly,
        DaysOfWeek =
        {
        SqlVmAutoBackupDayOfWeek.Monday,SqlVmAutoBackupDayOfWeek.Friday
        },
        FullBackupStartHour = 6,
        FullBackupWindowHours = 11,
        LogBackupFrequency = 10,
    },
    KeyVaultCredentialSettings = new SqlVmKeyVaultCredentialSettings()
    {
        IsEnabled = false,
    },
    ServerConfigurationsManagementSettings = new SqlServerConfigurationsManagementSettings()
    {
        SqlConnectivityUpdateSettings = new SqlConnectivityUpdateSettings()
        {
            ConnectivityType = SqlServerConnectivityType.Private,
            Port = 1433,
            SqlAuthUpdateUserName = "sqllogin",
            SqlAuthUpdatePassword = "<password>",
        },
        SqlWorkloadType = SqlWorkloadType.Oltp,
        SqlStorageUpdateSettings = new SqlStorageUpdateSettings()
        {
            DiskCount = 1,
            StartingDeviceId = 2,
            DiskConfigurationType = SqlVmDiskConfigurationType.New,
        },
        IsRServicesEnabled = false,
    },
};
ArmOperation<SqlVmResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sqlVmName, data);
SqlVmResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlVmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
