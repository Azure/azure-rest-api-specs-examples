import com.azure.resourcemanager.peering.models.PeeringLocationsKind;

/** Samples for PeeringLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListDirectPeeringLocations.json
     */
    /**
     * Sample code: List direct peering locations.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listDirectPeeringLocations(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peeringLocations().list(PeeringLocationsKind.DIRECT, null, com.azure.core.util.Context.NONE);
    }
}
