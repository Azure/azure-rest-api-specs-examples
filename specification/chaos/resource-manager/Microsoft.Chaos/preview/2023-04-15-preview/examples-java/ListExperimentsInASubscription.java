/** Samples for Experiments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListExperimentsInASubscription.json
     */
    /**
     * Sample code: List all Experiments in a subscription.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllExperimentsInASubscription(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.experiments().list(null, null, com.azure.core.util.Context.NONE);
    }
}
