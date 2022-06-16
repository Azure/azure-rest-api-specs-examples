import com.azure.core.util.Context;

/** Samples for Snapshots GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/GetInformationAboutASnapshot.json
     */
    /**
     * Sample code: Get information about a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .getByResourceGroupWithResponse("myResourceGroup", "mySnapshot", Context.NONE);
    }
}
