
/**
 * Samples for SnapshotPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/SnapshotPolicies_List.json
     */
    /**
     * Sample code: SnapshotPolicies_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().list("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
