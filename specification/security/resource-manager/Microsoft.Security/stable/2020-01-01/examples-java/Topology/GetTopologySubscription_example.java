
/**
 * Samples for Topology List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/
     * GetTopologySubscription_example.json
     */
    /**
     * Sample code: Get topology on a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getTopologyOnASubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.topologies().list(com.azure.core.util.Context.NONE);
    }
}
