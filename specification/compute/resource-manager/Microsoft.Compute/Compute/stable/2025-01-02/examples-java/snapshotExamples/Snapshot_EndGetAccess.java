
/**
 * Samples for Snapshots RevokeAccess.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_EndGetAccess.json
     */
    /**
     * Sample code: Revoke access to a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void revokeAccessToASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().revokeAccess("myResourceGroup", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
