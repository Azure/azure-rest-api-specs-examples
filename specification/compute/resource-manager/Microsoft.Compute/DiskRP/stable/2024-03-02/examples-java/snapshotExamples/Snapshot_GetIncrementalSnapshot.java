
/**
 * Samples for Snapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/snapshotExamples/
     * Snapshot_GetIncrementalSnapshot.json
     */
    /**
     * Sample code: Get information about an incremental snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAnIncrementalSnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots().getByResourceGroupWithResponse(
            "myResourceGroup", "myIncrementalSnapshot", com.azure.core.util.Context.NONE);
    }
}
