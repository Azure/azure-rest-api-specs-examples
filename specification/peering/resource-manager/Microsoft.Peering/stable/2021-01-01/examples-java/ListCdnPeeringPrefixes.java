/** Samples for CdnPeeringPrefixes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListCdnPeeringPrefixes.json
     */
    /**
     * Sample code: List all the cdn peering prefixes advertised at a particular peering location.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listAllTheCdnPeeringPrefixesAdvertisedAtAParticularPeeringLocation(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.cdnPeeringPrefixes().list("peeringLocation0", com.azure.core.util.Context.NONE);
    }
}
