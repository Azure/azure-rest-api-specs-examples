
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job/AutoMLJob/list.json
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
