
/**
 * Samples for Experiments ExecutionDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_ExecutionDetails.json
     */
    /**
     * Sample code: Get experiment execution details.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getExperimentExecutionDetails(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().executionDetailsWithResponse("exampleRG", "exampleExperiment",
            "f24500ad-744e-4a26-864b-b76199eac333", com.azure.core.util.Context.NONE);
    }
}
