
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/CommandJob/get.json
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
