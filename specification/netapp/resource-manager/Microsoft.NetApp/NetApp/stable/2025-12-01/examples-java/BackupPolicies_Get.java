
/**
 * Samples for BackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/BackupPolicies_Get.json
     */
    /**
     * Sample code: BackupPolicies_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupPoliciesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().getWithResponse("myRG", "account1", "backupPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
