
/**
 * Samples for SnapshotPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/SnapshotPolicies_Get.
     * json
     */
    /**
     * Sample code: SnapshotPolicies_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().getWithResponse("myRG", "account1", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
