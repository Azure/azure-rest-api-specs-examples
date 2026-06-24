
/**
 * Samples for RestorePoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/restorePointExamples/RestorePoint_Get_WithInstanceView.json
     */
    /**
     * Sample code: Get restore point with instance view.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getRestorePointWithInstanceView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().getWithResponse("myResourceGroup", "rpcName", "rpName", null,
            com.azure.core.util.Context.NONE);
    }
}
