/** Samples for Prefixes ListByPeeringService. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPrefixesByPeeringService.json
     */
    /**
     * Sample code: List all the prefixes associated with the peering service.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listAllThePrefixesAssociatedWithThePeeringService(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.prefixes().listByPeeringService("rgName", "peeringServiceName", null, com.azure.core.util.Context.NONE);
    }
}
