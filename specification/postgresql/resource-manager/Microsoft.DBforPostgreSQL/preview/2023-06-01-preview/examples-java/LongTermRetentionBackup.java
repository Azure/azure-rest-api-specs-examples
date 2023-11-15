import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupSettings;
import com.azure.resourcemanager.postgresqlflexibleserver.models.BackupStoreDetails;
import com.azure.resourcemanager.postgresqlflexibleserver.models.LtrBackupRequest;
import java.util.Arrays;

/** Samples for FlexibleServer StartLtrBackup. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/LongTermRetentionBackup.json
     */
    /**
     * Sample code: Sample_ExecuteBackup.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void sampleExecuteBackup(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .flexibleServers()
            .startLtrBackup(
                "rgLongTermRetention",
                "pgsqlltrtestserver",
                new LtrBackupRequest()
                    .withBackupSettings(new BackupSettings().withBackupName("backup1"))
                    .withTargetDetails(new BackupStoreDetails().withSasUriList(Arrays.asList("sasuri"))),
                com.azure.core.util.Context.NONE);
    }
}
