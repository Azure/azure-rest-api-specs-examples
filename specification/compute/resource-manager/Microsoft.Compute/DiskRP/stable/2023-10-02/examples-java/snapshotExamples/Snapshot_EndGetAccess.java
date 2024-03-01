
/**
 * Samples for Snapshots RevokeAccess.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/snapshotExamples/
     * Snapshot_EndGetAccess.json
     */
    /**
     * Sample code: Revoke access to a snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revokeAccessToASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getSnapshots().revokeAccess("myResourceGroup", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
