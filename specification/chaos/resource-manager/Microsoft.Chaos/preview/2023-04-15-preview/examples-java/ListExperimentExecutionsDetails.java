/** Samples for Experiments ListExecutionDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListExperimentExecutionsDetails.json
     */
    /**
     * Sample code: List experiment executions details.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listExperimentExecutionsDetails(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().listExecutionDetails("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
