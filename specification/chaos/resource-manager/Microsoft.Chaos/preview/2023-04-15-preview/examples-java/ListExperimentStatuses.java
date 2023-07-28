/** Samples for Experiments ListAllStatuses. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListExperimentStatuses.json
     */
    /**
     * Sample code: List all statuses of a Experiment.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllStatusesOfAExperiment(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().listAllStatuses("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
