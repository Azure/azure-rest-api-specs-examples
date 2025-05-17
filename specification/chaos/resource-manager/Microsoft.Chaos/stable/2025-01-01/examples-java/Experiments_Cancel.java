
/**
 * Samples for Experiments Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_Cancel.json
     */
    /**
     * Sample code: Cancel a running Experiment.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void cancelARunningExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().cancel("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
