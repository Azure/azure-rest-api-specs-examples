import com.azure.core.util.Context;

/** Samples for Disks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/DeleteAManagedDisk.json
     */
    /**
     * Sample code: Delete a managed disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAManagedDisk(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().delete("myResourceGroup", "myDisk", Context.NONE);
    }
}
