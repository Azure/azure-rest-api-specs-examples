import com.azure.core.util.Context;

/** Samples for Disks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/GetInformationAboutAManagedDisk.json
     */
    /**
     * Sample code: Get information about a managed disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .getByResourceGroupWithResponse("myResourceGroup", "myManagedDisk", Context.NONE);
    }
}
