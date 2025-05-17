
/**
 * Samples for Experiments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_Delete.json
     */
    /**
     * Sample code: Delete a Experiment in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAExperimentInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().delete("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
