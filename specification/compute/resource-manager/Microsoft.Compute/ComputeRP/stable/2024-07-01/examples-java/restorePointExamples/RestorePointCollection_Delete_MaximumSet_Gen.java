
/**
 * Samples for RestorePointCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * restorePointExamples/RestorePointCollection_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePointCollection_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restorePointCollectionDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePointCollections().delete("rgcompute",
            "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
