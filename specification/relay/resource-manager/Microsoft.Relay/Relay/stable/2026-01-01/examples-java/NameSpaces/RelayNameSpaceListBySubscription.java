
/**
 * Samples for Namespaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceListBySubscription.json
     */
    /**
     * Sample code: RelayNameSpaceListBySubscription.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceListBySubscription(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().list(com.azure.core.util.Context.NONE);
    }
}
