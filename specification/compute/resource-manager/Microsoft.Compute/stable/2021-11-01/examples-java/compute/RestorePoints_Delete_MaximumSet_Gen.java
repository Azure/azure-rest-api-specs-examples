import com.azure.core.util.Context;

/** Samples for RestorePoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/RestorePoints_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePoints_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restorePointsDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePoints()
            .delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", "a", Context.NONE);
    }
}
