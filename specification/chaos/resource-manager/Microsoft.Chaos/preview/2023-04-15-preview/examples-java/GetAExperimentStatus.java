/** Samples for Experiments GetStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetAExperimentStatus.json
     */
    /**
     * Sample code: Get the status of a Experiment.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getTheStatusOfAExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .getStatusWithResponse(
                "exampleRG",
                "exampleExperiment",
                "50734542-2e64-4e08-814c-cc0e7475f7e4",
                com.azure.core.util.Context.NONE);
    }
}
