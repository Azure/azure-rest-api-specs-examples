/** Samples for RestorePointCollections ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/restorePointExamples/RestorePointCollection_ListByResourceGroup.json
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
            .listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
