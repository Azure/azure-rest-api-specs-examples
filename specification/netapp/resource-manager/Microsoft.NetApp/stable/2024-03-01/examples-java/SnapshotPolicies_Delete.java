
/**
 * Samples for SnapshotPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/SnapshotPolicies_Delete.json
     */
    /**
     * Sample code: SnapshotPolicies_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().delete("resourceGroup", "accountName", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
