import com.azure.core.util.Context;

/** Samples for Snapshots ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/snapshotExamples/Snapshot_ListByResourceGroup.json
     */
    /**
     * Sample code: List all snapshots in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllSnapshotsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
