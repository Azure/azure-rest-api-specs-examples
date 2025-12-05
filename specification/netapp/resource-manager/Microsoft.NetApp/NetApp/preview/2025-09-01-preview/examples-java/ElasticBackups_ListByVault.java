
/**
 * Samples for ElasticBackups ListByVault.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackups_ListByVault.json
     */
    /**
     * Sample code: ElasticBackups_ListByVault.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupsListByVault(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackups().listByVault("myRG", "account1", "backupVault1", com.azure.core.util.Context.NONE);
    }
}
