/** Samples for PeeringServiceLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringServiceLocationsBySubscription.json
     */
    /**
     * Sample code: List peering service locations.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringServiceLocations(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringServiceLocations().list(null, com.azure.core.util.Context.NONE);
    }
}
