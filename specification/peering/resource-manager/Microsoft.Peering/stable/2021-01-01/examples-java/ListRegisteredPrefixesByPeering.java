/** Samples for RegisteredPrefixes ListByPeering. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListRegisteredPrefixesByPeering.json
     */
    /**
     * Sample code: List all the registered prefixes associated with the peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listAllTheRegisteredPrefixesAssociatedWithThePeering(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.registeredPrefixes().listByPeering("rgName", "peeringName", com.azure.core.util.Context.NONE);
    }
}
