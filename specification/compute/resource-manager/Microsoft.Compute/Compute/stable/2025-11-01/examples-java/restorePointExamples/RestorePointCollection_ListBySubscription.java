
/**
 * Samples for RestorePointCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePointCollection_ListBySubscription.json
     */
    /**
     * Sample code: Gets the list of restore point collections in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getsTheListOfRestorePointCollectionsInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().list(com.azure.core.util.Context.NONE);
    }
}
