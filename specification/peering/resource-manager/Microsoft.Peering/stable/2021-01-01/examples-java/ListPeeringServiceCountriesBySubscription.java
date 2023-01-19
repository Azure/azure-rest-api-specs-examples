/** Samples for PeeringServiceCountries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringServiceCountriesBySubscription.json
     */
    /**
     * Sample code: List peering service countries.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringServiceCountries(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringServiceCountries().list(com.azure.core.util.Context.NONE);
    }
}
