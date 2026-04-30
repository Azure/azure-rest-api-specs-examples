
/**
 * Samples for RestorePointCollections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePointCollection_ListByResourceGroup.json
     */
    /**
     * Sample code: Gets the list of restore point collections in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getsTheListOfRestorePointCollectionsInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().listByResourceGroup("myResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
