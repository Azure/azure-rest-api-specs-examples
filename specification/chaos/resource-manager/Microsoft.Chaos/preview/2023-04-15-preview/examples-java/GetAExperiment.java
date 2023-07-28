/** Samples for Experiments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetAExperiment.json
     */
    /**
     * Sample code: Get a Experiment in a resource group.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getAExperimentInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .getByResourceGroupWithResponse("exampleRG", "exampleExperiment", com.azure.core.util.Context.NONE);
    }
}
