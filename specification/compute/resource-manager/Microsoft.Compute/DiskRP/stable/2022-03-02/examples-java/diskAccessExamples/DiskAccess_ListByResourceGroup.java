import com.azure.core.util.Context;

/** Samples for DiskAccesses ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskAccessExamples/DiskAccess_ListByResourceGroup.json
     */
    /**
     * Sample code: List all disk access resources in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllDiskAccessResourcesInAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
