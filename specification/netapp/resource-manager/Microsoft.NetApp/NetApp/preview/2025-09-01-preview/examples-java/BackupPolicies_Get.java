
/**
 * Samples for BackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupPolicies_Get.json
     */
    /**
     * Sample code: Backups_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().getWithResponse("myRG", "account1", "backupPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
