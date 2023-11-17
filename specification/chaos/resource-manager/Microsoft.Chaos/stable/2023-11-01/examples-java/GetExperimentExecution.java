/** Samples for Experiments GetExecution. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetExperimentExecution.json
     */
    /**
     * Sample code: Get the execution of a Experiment.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getTheExecutionOfAExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .getExecutionWithResponse(
                "exampleRG",
                "exampleExperiment",
                "f24500ad-744e-4a26-864b-b76199eac333",
                com.azure.core.util.Context.NONE);
    }
}
