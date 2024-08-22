
/**
 * Samples for Schedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Schedule/get.json
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
