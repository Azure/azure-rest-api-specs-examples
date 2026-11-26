
/**
 * Samples for Schedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Schedule/get.json
     */
    /**
     * Sample code: Get Schedule.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getSchedule(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.schedules().getWithResponse("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
