/** Samples for PeeringServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/GetPeeringService.json
     */
    /**
     * Sample code: Get a peering service.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void getAPeeringService(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .peeringServices()
            .getByResourceGroupWithResponse("rgName", "peeringServiceName", com.azure.core.util.Context.NONE);
    }
}
