import com.azure.core.util.Context;

/** Samples for RestorePointCollections ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointCollectionsInAResourceGroup.json
     */
    /**
     * Sample code: Gets the list of restore point collections in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheListOfRestorePointCollectionsInAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
