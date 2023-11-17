/** Samples for Experiments Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/StartExperiment.json
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
