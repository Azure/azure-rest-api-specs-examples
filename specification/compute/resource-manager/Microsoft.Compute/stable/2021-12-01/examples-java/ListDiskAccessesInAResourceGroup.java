import com.azure.core.util.Context;

/** Samples for DiskAccesses ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/ListDiskAccessesInAResourceGroup.json
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
