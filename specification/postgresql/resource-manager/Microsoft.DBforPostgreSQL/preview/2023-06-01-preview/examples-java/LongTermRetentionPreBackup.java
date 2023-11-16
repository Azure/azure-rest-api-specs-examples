import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupSettings;
import com.azure.resourcemanager.postgresqlflexibleserver.models.LtrPreBackupRequest;

/** Samples for FlexibleServer TriggerLtrPreBackup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/LongTermRetentionPreBackup.json
     */
    /**
     * Sample code: Sample_Prebackup.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void samplePrebackup(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .flexibleServers()
            .triggerLtrPreBackupWithResponse(
                "rgLongTermRetention",
                "pgsqlltrtestserver",
                new LtrPreBackupRequest().withBackupSettings(new BackupSettings().withBackupName("backup1")),
                com.azure.core.util.Context.NONE);
    }
}
