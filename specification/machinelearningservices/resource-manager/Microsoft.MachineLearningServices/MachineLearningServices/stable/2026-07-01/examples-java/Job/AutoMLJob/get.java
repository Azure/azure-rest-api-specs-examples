
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job/AutoMLJob/get.json
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
