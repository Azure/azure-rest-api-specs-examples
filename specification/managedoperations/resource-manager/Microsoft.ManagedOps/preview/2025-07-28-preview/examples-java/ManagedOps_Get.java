
/**
 * Samples for ManagedOps Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-28-preview/ManagedOps_Get.json
     */
    /**
     * Sample code: ManagedOps_Get.
     * 
     * @param manager Entry point to ManagedOpsManager.
     */
    public static void managedOpsGet(com.azure.resourcemanager.managedops.ManagedOpsManager manager) {
        manager.managedOps().getWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
