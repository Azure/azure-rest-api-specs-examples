
import com.azure.resourcemanager.appservice.fluent.models.BackupRequestInner;
import com.azure.resourcemanager.appservice.models.BackupSchedule;
import com.azure.resourcemanager.appservice.models.DatabaseBackupSetting;
import com.azure.resourcemanager.appservice.models.DatabaseType;
import com.azure.resourcemanager.appservice.models.FrequencyUnit;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for WebApps Backup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/BackupWebApp.json
     */
    /**
     * Sample code: Backup web app.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void backupWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().backupWithResponse("testrg123", "sitef6141",
            new BackupRequestInner().withBackupName("abcdwe").withEnabled(true)
                .withStorageAccountUrl(
                    "DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>")
                .withBackupSchedule(new BackupSchedule().withFrequencyInterval(7).withFrequencyUnit(FrequencyUnit.DAY)
                    .withKeepAtLeastOneBackup(true).withRetentionPeriodInDays(30)
                    .withStartTime(OffsetDateTime.parse("2022-09-02T17:33:11.641Z")))
                .withDatabases(Arrays.asList(
                    new DatabaseBackupSetting().withDatabaseType(DatabaseType.SQL_AZURE).withName("backenddb")
                        .withConnectionStringName("backend").withConnectionString(
                            "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"),
                    new DatabaseBackupSetting().withDatabaseType(DatabaseType.SQL_AZURE).withName("statsdb")
                        .withConnectionStringName("stats").withConnectionString(
                            "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]"))),
            com.azure.core.util.Context.NONE);
    }
}
