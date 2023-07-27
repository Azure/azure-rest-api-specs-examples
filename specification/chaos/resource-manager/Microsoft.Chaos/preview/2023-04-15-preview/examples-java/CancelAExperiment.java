/** Samples for Experiments Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/CancelAExperiment.json
     */
    /**
     * Sample code: Cancel a running Experiment.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void cancelARunningExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().cancelWithResponse("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
