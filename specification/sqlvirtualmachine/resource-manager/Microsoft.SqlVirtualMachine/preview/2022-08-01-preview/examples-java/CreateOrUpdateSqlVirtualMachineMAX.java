
import com.azure.resourcemanager.sqlvirtualmachine.models.AadAuthenticationSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.AdditionalFeaturesServerConfigurations;
import com.azure.resourcemanager.sqlvirtualmachine.models.AssessmentDayOfWeek;
import com.azure.resourcemanager.sqlvirtualmachine.models.AssessmentSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.AutoBackupSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.AutoPatchingSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.BackupScheduleType;
import com.azure.resourcemanager.sqlvirtualmachine.models.ConnectivityType;
import com.azure.resourcemanager.sqlvirtualmachine.models.DayOfWeek;
import com.azure.resourcemanager.sqlvirtualmachine.models.DiskConfigurationType;
import com.azure.resourcemanager.sqlvirtualmachine.models.FullBackupFrequencyType;
import com.azure.resourcemanager.sqlvirtualmachine.models.KeyVaultCredentialSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.LeastPrivilegeMode;
import com.azure.resourcemanager.sqlvirtualmachine.models.Schedule;
import com.azure.resourcemanager.sqlvirtualmachine.models.ServerConfigurationsManagementSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlConnectivityUpdateSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlImageSku;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlInstanceSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlManagementMode;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlServerLicenseType;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlStorageUpdateSettings;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlWorkloadType;
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlWorkloadTypeUpdateSettings;

/**
 * Samples for SqlVirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateSqlVirtualMachineMAX.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine with max parameters.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesASQLVirtualMachineWithMaxParameters(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().define("testvm").withRegion("northeurope").withExistingResourceGroup("testrg")
            .withVirtualMachineResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm")
            .withSqlServerLicenseType(SqlServerLicenseType.PAYG).withSqlManagement(SqlManagementMode.FULL)
            .withLeastPrivilegeMode(LeastPrivilegeMode.ENABLED).withSqlImageSku(SqlImageSku.ENTERPRISE)
            .withAutoPatchingSettings(new AutoPatchingSettings().withEnable(true).withDayOfWeek(DayOfWeek.SUNDAY)
                .withMaintenanceWindowStartingHour(2).withMaintenanceWindowDuration(60))
            .withAutoBackupSettings(new AutoBackupSettings().withEnable(true).withEnableEncryption(true)
                .withRetentionPeriod(17).withStorageAccountUrl("https://teststorage.blob.core.windows.net/")
                .withStorageContainerName("testcontainer").withStorageAccessKey("fakeTokenPlaceholder")
                .withPassword("fakeTokenPlaceholder").withBackupSystemDbs(true)
                .withBackupScheduleType(BackupScheduleType.MANUAL)
                .withFullBackupFrequency(FullBackupFrequencyType.DAILY).withFullBackupStartTime(6)
                .withFullBackupWindowHours(11).withLogBackupFrequency(10))
            .withKeyVaultCredentialSettings(new KeyVaultCredentialSettings().withEnable(false))
            .withServerConfigurationsManagementSettings(new ServerConfigurationsManagementSettings()
                .withSqlConnectivityUpdateSettings(
                    new SqlConnectivityUpdateSettings().withConnectivityType(ConnectivityType.PRIVATE).withPort(1433)
                        .withSqlAuthUpdateUsername("sqllogin").withSqlAuthUpdatePassword("fakeTokenPlaceholder"))
                .withSqlWorkloadTypeUpdateSettings(
                    new SqlWorkloadTypeUpdateSettings().withSqlWorkloadType(SqlWorkloadType.OLTP))
                .withSqlStorageUpdateSettings(new SqlStorageUpdateSettings().withDiskCount(1).withStartingDeviceId(2)
                    .withDiskConfigurationType(DiskConfigurationType.NEW))
                .withAdditionalFeaturesServerConfigurations(
                    new AdditionalFeaturesServerConfigurations().withIsRServicesEnabled(false))
                .withSqlInstanceSettings(new SqlInstanceSettings().withCollation("SQL_Latin1_General_CP1_CI_AS")
                    .withMaxDop(8).withIsOptimizeForAdHocWorkloadsEnabled(true).withMinServerMemoryMB(0)
                    .withMaxServerMemoryMB(128).withIsLpimEnabled(true).withIsIfiEnabled(true))
                .withAzureAdAuthenticationSettings(
                    new AadAuthenticationSettings().withClientId("11111111-2222-3333-4444-555555555555")))
            .withAssessmentSettings(new AssessmentSettings().withEnable(true).withRunImmediately(true)
                .withSchedule(new Schedule().withEnable(true).withWeeklyInterval(1)
                    .withDayOfWeek(AssessmentDayOfWeek.SUNDAY).withStartTime("23:17")))
            .withEnableAutomaticUpgrade(true).create();
    }
}
