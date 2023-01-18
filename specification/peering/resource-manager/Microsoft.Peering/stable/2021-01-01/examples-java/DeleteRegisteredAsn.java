/** Samples for RegisteredAsns Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/DeleteRegisteredAsn.json
     */
    /**
     * Sample code: Deletes a registered ASN associated with the peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void deletesARegisteredASNAssociatedWithThePeering(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .registeredAsns()
            .deleteWithResponse("rgName", "peeringName", "registeredAsnName", com.azure.core.util.Context.NONE);
    }
}
