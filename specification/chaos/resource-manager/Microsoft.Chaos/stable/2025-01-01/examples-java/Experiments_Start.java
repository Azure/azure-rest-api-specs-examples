
/**
 * Samples for Experiments Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_Start.json
     */
    /**
     * Sample code: Start a Experiment.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void startAExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().start("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
