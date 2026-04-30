
/**
 * Samples for RestorePointCollections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
     */
    /**
     * Sample code: Get a restore point collection, including the restore points contained in the restore point
     * collection.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().getByResourceGroupWithResponse("myResourceGroup",
            "rpcName", null, com.azure.core.util.Context.NONE);
    }
}
