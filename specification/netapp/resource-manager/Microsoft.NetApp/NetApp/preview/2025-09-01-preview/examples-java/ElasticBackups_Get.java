
/**
 * Samples for ElasticBackups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackups_Get.json
     */
    /**
     * Sample code: ElasticBackups_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackups().getWithResponse("myRG", "account1", "backupVault1", "backup1",
            com.azure.core.util.Context.NONE);
    }
}
