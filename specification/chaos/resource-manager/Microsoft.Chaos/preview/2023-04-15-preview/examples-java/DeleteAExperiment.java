/** Samples for Experiments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/DeleteAExperiment.json
     */
    /**
     * Sample code: Delete a Experiment in a resource group.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAExperimentInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .deleteByResourceGroupWithResponse("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
