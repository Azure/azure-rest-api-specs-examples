/** Samples for PeeringServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/DeletePeeringService.json
     */
    /**
     * Sample code: Delete a peering service.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void deleteAPeeringService(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peeringServices()
            .deleteByResourceGroupWithResponse("rgName", "peeringServiceName", com.azure.core.util.Context.NONE);
    }
}
