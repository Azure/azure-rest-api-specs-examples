/** Samples for RegisteredAsns CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/CreateRegisteredAsn.json
     */
    /**
     * Sample code: Create or update a registered ASN for the peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void createOrUpdateARegisteredASNForThePeering(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .registeredAsns()
            .define("registeredAsnName")
            .withExistingPeering("rgName", "peeringName")
            .withAsn(65000)
            .create();
    }
}
