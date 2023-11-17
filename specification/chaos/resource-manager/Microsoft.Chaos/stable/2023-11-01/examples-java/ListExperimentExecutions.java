/** Samples for Experiments ListAllExecutions. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListExperimentExecutions.json
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
