
/**
 * Samples for RestorePointCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * restorePointExamples/RestorePointCollection_ListBySubscription.json
     */
    /**
     * Sample code: Gets the list of restore point collections in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsTheListOfRestorePointCollectionsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePointCollections()
            .list(com.azure.core.util.Context.NONE);
    }
}
