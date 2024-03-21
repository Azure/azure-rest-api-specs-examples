
/**
 * Samples for SnapshotPolicies ListVolumes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/SnapshotPolicies_ListVolumes.
     * json
     */
    /**
     * Sample code: SnapshotPolicies_ListVolumes.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesListVolumes(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().listVolumesWithResponse("myRG", "account1", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
