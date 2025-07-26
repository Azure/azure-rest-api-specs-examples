
/**
 * Samples for DiskAccesses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskAccessExamples/
     * DiskAccess_ListByResourceGroup.json
     */
    /**
     * Sample code: list all disk access resources in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllDiskAccessResourcesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskAccesses().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
