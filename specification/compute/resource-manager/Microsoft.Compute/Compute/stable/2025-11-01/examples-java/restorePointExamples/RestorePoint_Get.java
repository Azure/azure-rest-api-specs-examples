
/**
 * Samples for RestorePoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/restorePointExamples/RestorePoint_Get.json
     */
    /**
     * Sample code: Get a restore point.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getARestorePoint(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getRestorePoints().getWithResponse("myResourceGroup", "rpcName", "rpName", null,
            com.azure.core.util.Context.NONE);
    }
}
