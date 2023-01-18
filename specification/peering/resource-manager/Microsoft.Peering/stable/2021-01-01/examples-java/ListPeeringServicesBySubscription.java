/** Samples for PeeringServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringServicesBySubscription.json
     */
    /**
     * Sample code: List peering services in a subscription.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringServicesInASubscription(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringServices().list(com.azure.core.util.Context.NONE);
    }
}
