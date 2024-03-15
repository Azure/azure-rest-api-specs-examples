
/**
 * Samples for Experiments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/GetExperiment.json
     */
    /**
     * Sample code: Get a Experiment in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAExperimentInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().getByResourceGroupWithResponse("exampleRG", "exampleExperiment",
            com.azure.core.util.Context.NONE);
    }
}
