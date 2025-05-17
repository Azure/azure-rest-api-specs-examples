
/**
 * Samples for Experiments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Experiments_ListAll.json
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
