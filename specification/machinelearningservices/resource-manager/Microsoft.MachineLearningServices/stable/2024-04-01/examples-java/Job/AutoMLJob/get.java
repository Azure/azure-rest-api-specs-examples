
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/AutoMLJob/get.json
     */
    /**
     * Sample code: Get AutoML Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getAutoMLJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().getWithResponse("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
