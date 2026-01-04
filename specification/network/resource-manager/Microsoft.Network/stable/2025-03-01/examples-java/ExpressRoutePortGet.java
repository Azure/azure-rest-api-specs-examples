
/**
 * Samples for ExpressRoutePorts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRoutePortGet.json
     */
    /**
     * Sample code: ExpressRoutePortGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void expressRoutePortGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRoutePorts().getByResourceGroupWithResponse("rg1",
            "portName", com.azure.core.util.Context.NONE);
    }
}
