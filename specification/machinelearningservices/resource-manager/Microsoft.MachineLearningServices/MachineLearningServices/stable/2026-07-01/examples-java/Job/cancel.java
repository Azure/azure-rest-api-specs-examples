
/**
 * Samples for Jobs Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Job/cancel.json
     */
    /**
     * Sample code: Cancel Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void cancelJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.jobs().cancel("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
