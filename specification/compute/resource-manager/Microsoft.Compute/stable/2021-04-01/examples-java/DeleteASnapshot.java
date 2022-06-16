import com.azure.core.util.Context;

/** Samples for Snapshots Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/DeleteASnapshot.json
     */
    /**
     * Sample code: Delete a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .delete("myResourceGroup", "mySnapshot", Context.NONE);
    }
}
