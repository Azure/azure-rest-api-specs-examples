/** Samples for PeeringServiceProviders List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringServiceProviders.json
     */
    /**
     * Sample code: List peering service providers.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringServiceProviders(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringServiceProviders().list(com.azure.core.util.Context.NONE);
    }
}
