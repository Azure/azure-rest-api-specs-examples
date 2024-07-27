
/**
 * Samples for Snapshots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/
     * Snapshot_ListByResourceGroup.json
     */
    /**
     * Sample code: List all snapshots in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllSnapshotsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
