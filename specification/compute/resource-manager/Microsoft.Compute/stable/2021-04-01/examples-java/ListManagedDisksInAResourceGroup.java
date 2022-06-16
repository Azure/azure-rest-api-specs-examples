import com.azure.core.util.Context;

/** Samples for Disks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/ListManagedDisksInAResourceGroup.json
     */
    /**
     * Sample code: List all managed disks in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllManagedDisksInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
