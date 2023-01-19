/** Samples for Peerings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/DeletePeering.json
     */
    /**
     * Sample code: Delete a peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void deleteAPeering(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peerings().deleteByResourceGroupWithResponse("rgName", "peeringName", com.azure.core.util.Context.NONE);
    }
}
