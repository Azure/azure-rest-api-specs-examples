import com.azure.core.util.Context;

/** Samples for SnapshotPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/SnapshotPolicies_Get.json
     */
    /**
     * Sample code: SnapshotPolicies_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().getWithResponse("myRG", "account1", "snapshotPolicyName", Context.NONE);
    }
}
