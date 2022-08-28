import com.azure.core.util.Context;

/** Samples for RestorePointCollections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/restorePointExamples/RestorePointCollections_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePointCollections_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restorePointCollectionsDeleteMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getRestorePointCollections()
            .delete("rgcompute", "aaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
