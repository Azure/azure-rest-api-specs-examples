
/**
 * Samples for Snapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/
     * Snapshot_Delete.json
     */
    /**
     * Sample code: Delete a snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots().delete("myResourceGroup", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
