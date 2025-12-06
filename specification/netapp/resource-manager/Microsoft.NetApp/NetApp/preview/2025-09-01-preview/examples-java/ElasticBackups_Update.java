
import com.azure.resourcemanager.netapp.models.ElasticBackup;

/**
 * Samples for ElasticBackups Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackups_Update.json
     */
    /**
     * Sample code: ElasticBackups_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        ElasticBackup resource = manager.elasticBackups()
            .getWithResponse("myRG", "account1", "backupVault1", "backup1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
