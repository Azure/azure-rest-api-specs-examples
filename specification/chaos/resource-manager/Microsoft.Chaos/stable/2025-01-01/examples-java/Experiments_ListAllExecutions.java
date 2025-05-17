
/**
 * Samples for Experiments ListAllExecutions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_ListAllExecutions.json
     */
    /**
     * Sample code: List all executions of an Experiment.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllExecutionsOfAnExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().listAllExecutions("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
