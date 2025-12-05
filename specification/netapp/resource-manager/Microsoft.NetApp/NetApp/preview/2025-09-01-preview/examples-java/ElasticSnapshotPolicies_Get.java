
/**
 * Samples for ElasticSnapshotPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticSnapshotPolicies_Get.json
     */
    /**
     * Sample code: ElasticSnapshotPolicies_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticSnapshotPoliciesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshotPolicies().getWithResponse("myRG", "account1", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
