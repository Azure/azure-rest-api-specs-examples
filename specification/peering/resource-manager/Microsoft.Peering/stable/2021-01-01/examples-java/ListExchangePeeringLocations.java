import com.azure.resourcemanager.peering.models.PeeringLocationsKind;

/** Samples for PeeringLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListExchangePeeringLocations.json
     */
    /**
     * Sample code: List exchange peering locations.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listExchangePeeringLocations(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringLocations().list(PeeringLocationsKind.EXCHANGE, null, com.azure.core.util.Context.NONE);
    }
}
