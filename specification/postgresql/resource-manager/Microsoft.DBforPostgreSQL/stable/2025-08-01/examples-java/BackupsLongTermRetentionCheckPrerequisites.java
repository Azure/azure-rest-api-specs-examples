
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupSettings;
import com.azure.resourcemanager.postgresqlflexibleserver.models.LtrPreBackupRequest;

/**
 * Samples for BackupsLongTermRetention CheckPrerequisites.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsLongTermRetentionCheckPrerequisites.json
     */
    /**
     * Sample code: Perform all checks required for a long term retention backup operation to succeed.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void performAllChecksRequiredForALongTermRetentionBackupOperationToSucceed(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsLongTermRetentions().checkPrerequisitesWithResponse("exampleresourcegroup", "exampleserver",
            new LtrPreBackupRequest().withBackupSettings(new BackupSettings().withBackupName("exampleltrbackup")),
            com.azure.core.util.Context.NONE);
    }
}
