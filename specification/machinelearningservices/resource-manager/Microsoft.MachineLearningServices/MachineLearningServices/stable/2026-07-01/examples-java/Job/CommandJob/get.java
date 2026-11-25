
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job/CommandJob/get.json
     */
    /**
     * Sample code: Get Command Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getCommandJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().getWithResponse("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
