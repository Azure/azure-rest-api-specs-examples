import com.azure.core.util.Context;

/** Samples for RestorePointCollections GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
     */
    /**
     * Sample code: Get a restore point collection, including the restore points contained in the restore point
     * collection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .getByResourceGroupWithResponse("myResourceGroup", "rpcName", null, Context.NONE);
    }
}
