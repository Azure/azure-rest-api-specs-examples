
/**
 * Samples for ElasticBackupPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupPolicies_Get.json
     */
    /**
     * Sample code: ElasticBackupPolicies_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupPoliciesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupPolicies().getWithResponse("myRG", "account1", "backupPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
