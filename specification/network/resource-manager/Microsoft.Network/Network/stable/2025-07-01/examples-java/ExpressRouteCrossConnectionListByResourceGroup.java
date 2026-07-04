
/**
 * Samples for ExpressRouteCrossConnections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionListByResourceGroup.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteCrossConnectionListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().listByResourceGroup("CrossConnection-SiliconValley",
            com.azure.core.util.Context.NONE);
    }
}
