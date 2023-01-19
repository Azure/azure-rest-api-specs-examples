/** Samples for Peerings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringsBySubscription.json
     */
    /**
     * Sample code: List peerings in a subscription.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringsInASubscription(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peerings().list(com.azure.core.util.Context.NONE);
    }
}
