
/**
 * Samples for StandbyVirtualMachinePoolRuntimeViews Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/StandbyVirtualMachinePoolRuntimeViews_Get.json
     */
    /**
     * Sample code: StandbyVirtualMachinePoolRuntimeViews_Get.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyVirtualMachinePoolRuntimeViewsGet(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePoolRuntimeViews().getWithResponse("rgstandbypool", "pool", "latest",
            com.azure.core.util.Context.NONE);
    }
}
