
/**
 * Samples for ExpressRouteCrossConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionList.json
     */
    /**
     * Sample code: ExpressRouteCrossConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteCrossConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().list(null, com.azure.core.util.Context.NONE);
    }
}
