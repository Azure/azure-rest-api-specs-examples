
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/PipelineJob/list.json
     */
    /**
     * Sample code: List Pipeline Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listPipelineJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, "string", "string", null, null,
            com.azure.core.util.Context.NONE);
    }
}
