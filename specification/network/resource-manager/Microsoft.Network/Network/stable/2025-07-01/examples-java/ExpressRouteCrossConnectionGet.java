
/**
 * Samples for ExpressRouteCrossConnections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionGet.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCrossConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().getByResourceGroupWithResponse(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", com.azure.core.util.Context.NONE);
    }
}
