
/**
 * Samples for RestorePoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * restorePointExamples/RestorePoint_Get_WithInstanceView.json
     */
    /**
     * Sample code: Get restore point with instance view.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRestorePointWithInstanceView(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getRestorePoints().getWithResponse("myResourceGroup",
            "rpcName", "rpName", null, com.azure.core.util.Context.NONE);
    }
}
