
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupSettings;
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupStoreDetails;
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupsLongTermRetentionRequest;
import java.util.Arrays;

/**
 * Samples for BackupsLongTermRetention Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * BackupsLongTermRetentionStart.json
     */
    /**
     * Sample code: Initiate a long term retention backup.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        initiateALongTermRetentionBackup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.backupsLongTermRetentions().start("exampleresourcegroup", "exampleserver",
            new BackupsLongTermRetentionRequest()
                .withBackupSettings(new BackupSettings().withBackupName("exampleltrbackup"))
                .withTargetDetails(new BackupStoreDetails().withSasUriList(Arrays.asList("sasuri"))),
            com.azure.core.util.Context.NONE);
    }
}
