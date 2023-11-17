/** Samples for Experiments Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/CancelExperiment.json
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
