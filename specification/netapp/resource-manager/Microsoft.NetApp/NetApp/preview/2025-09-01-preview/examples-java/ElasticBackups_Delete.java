
/**
 * Samples for ElasticBackups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackups_Delete.json
     */
    /**
     * Sample code: ElasticBackups_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackups().delete("resourceGroup", "account1", "backupVault1", "backup1",
            com.azure.core.util.Context.NONE);
    }
}
