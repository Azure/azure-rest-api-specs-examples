import com.azure.core.util.Context;

/** Samples for RestorePointCollections GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/GetRestorePointCollection.json
     */
    /**
     * Sample code: Get a restore point collection (but not the restore points contained in the restore point
     * collection).
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getARestorePointCollectionButNotTheRestorePointsContainedInTheRestorePointCollection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .getByResourceGroupWithResponse("myResourceGroup", "myRpc", null, Context.NONE);
    }
}
