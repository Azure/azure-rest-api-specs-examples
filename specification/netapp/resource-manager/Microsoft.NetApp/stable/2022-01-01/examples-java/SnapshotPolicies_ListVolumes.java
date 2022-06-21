import com.azure.core.util.Context;

/** Samples for SnapshotPolicies ListVolumes. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/SnapshotPolicies_ListVolumes.json
     */
    /**
     * Sample code: SnapshotPolicies_ListVolumes.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesListVolumes(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshotPolicies().listVolumesWithResponse("myRG", "account1", "snapshotPolicyName", Context.NONE);
    }
}
