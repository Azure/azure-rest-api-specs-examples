
/**
 * Samples for RestorePointCollections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePointCollection_Get.json
     */
    /**
     * Sample code: Get a restore point collection (but not the restore points contained in the restore point
     * collection).
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePointCollections().getByResourceGroupWithResponse("myResourceGroup", "myRpc",
            null, com.azure.core.util.Context.NONE);
    }
}
