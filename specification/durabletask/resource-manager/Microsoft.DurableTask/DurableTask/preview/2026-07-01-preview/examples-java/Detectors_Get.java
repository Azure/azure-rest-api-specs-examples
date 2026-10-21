
/**
 * Samples for Detectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Detectors_Get.json
     */
    /**
     * Sample code: Detectors_Get.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void detectorsGet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.detectors().getWithResponse("rgopenapi", "testscheduler", "cpuhealth",
            com.azure.core.util.Context.NONE);
    }
}
