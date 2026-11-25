
/**
 * Samples for Jobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job/CommandJob/list.json
     */
    /**
     * Sample code: List Command Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listCommandJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().list("test-rg", "my-aml-workspace", null, "string", "string", null, null,
            com.azure.core.util.Context.NONE);
    }
}
