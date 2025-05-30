
/**
 * Samples for RestorePoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * restorePointExamples/RestorePoint_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RestorePoint_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restorePointDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePoints().delete("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaa", "a", com.azure.core.util.Context.NONE);
    }
}
