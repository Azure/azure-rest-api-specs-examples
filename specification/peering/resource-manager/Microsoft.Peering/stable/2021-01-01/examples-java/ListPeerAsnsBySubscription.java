/** Samples for PeerAsns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeerAsnsBySubscription.json
     */
    /**
     * Sample code: List peer ASNs in a subscription.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeerASNsInASubscription(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peerAsns().list(com.azure.core.util.Context.NONE);
    }
}
