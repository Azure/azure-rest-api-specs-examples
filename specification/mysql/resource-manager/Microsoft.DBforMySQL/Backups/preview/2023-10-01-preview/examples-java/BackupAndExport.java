
import com.azure.resourcemanager.mysqlflexibleserver.models.BackupAndExportRequest;
import com.azure.resourcemanager.mysqlflexibleserver.models.BackupSettings;
import com.azure.resourcemanager.mysqlflexibleserver.models.FullBackupStoreDetails;
import java.util.Arrays;

/**
 * Samples for BackupAndExport Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Backups/preview/2023-10-01-preview/examples/
     * BackupAndExport.json
     */
    /**
     * Sample code: Create and Export Backup.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void createAndExportBackup(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.backupAndExports().create("TestGroup", "mysqltestserver",
            new BackupAndExportRequest().withBackupSettings(new BackupSettings().withBackupName("customer-backup-name"))
                .withTargetDetails(new FullBackupStoreDetails().withSasUriList(Arrays.asList("sasuri1", "sasuri2"))),
            com.azure.core.util.Context.NONE);
    }
}
