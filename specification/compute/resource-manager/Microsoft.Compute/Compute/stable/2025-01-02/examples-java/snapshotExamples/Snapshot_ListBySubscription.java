
/**
 * Samples for Snapshots List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_ListBySubscription.json
     */
    /**
     * Sample code: List all snapshots in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllSnapshotsInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().list(com.azure.core.util.Context.NONE);
    }
}
