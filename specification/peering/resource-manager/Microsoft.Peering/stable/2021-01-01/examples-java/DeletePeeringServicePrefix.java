/** Samples for Prefixes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/DeletePeeringServicePrefix.json
     */
    /**
     * Sample code: Delete a prefix associated with the peering service.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void deleteAPrefixAssociatedWithThePeeringService(
        com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .prefixes()
            .deleteWithResponse(
                "rgName", "peeringServiceName", "peeringServicePrefixName", com.azure.core.util.Context.NONE);
    }
}
