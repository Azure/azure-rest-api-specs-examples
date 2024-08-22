
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/AutoMLJob/list.json
     */
    /**
     * Sample code: List AutoML Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listAutoMLJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
