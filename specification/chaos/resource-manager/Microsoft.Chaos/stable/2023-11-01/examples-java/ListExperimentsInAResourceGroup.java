/** Samples for Experiments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListExperimentsInAResourceGroup.json
     */
    /**
     * Sample code: List all Experiments in a resource group.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllExperimentsInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().listByResourceGroup("exampleRG", null, null, com.azure.core.util.Context.NONE);
    }
}
