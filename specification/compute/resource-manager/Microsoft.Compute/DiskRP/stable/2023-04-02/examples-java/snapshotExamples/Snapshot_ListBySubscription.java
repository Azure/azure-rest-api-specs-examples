
/**
 * Samples for Snapshots List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/snapshotExamples/
     * Snapshot_ListBySubscription.json
     */
    /**
     * Sample code: List all snapshots in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllSnapshotsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots().list(com.azure.core.util.Context.NONE);
    }
}
