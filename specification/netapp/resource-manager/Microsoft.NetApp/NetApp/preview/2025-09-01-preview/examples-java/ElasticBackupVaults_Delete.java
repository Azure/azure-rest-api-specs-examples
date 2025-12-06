
/**
 * Samples for ElasticBackupVaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupVaults_Delete.json
     */
    /**
     * Sample code: ElasticBackupVaults_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupVaultsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupVaults().delete("resourceGroup", "account1", "backupVault1",
            com.azure.core.util.Context.NONE);
    }
}
