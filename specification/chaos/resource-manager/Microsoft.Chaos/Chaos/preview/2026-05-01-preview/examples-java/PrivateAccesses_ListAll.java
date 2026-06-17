
/**
 * Samples for PrivateAccesses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_ListAll.json
     */
    /**
     * Sample code: List all private accesses in a subscription.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllPrivateAccessesInASubscription(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().list(null, com.azure.core.util.Context.NONE);
    }
}
