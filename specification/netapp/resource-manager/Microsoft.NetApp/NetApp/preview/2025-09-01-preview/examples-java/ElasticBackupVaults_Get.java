
/**
 * Samples for ElasticBackupVaults Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupVaults_Get.json
     */
    /**
     * Sample code: ElasticBackupVaults_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupVaultsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupVaults().getWithResponse("myRG", "account1", "backupVault1",
            com.azure.core.util.Context.NONE);
    }
}
