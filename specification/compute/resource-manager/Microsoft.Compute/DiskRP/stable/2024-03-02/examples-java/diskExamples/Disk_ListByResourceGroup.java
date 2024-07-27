
/**
 * Samples for Disks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_ListByResourceGroup.json
     */
    /**
     * Sample code: List all managed disks in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllManagedDisksInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
