
/**
 * Samples for RestorePointCollections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * restorePointExamples/RestorePointCollection_Get.json
     */
    /**
     * Sample code: Get a restore point collection (but not the restore points contained in the restore point
     * collection).
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePointCollections()
            .getByResourceGroupWithResponse("myResourceGroup", "myRpc", null, com.azure.core.util.Context.NONE);
    }
}
