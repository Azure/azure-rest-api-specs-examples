/** Samples for Peerings GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/GetPeering.json
     */
    /**
     * Sample code: Get a peering.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void getAPeering(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peerings().getByResourceGroupWithResponse("rgName", "peeringName", com.azure.core.util.Context.NONE);
    }
}
