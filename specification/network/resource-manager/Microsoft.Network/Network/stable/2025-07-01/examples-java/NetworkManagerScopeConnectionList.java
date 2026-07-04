
/**
 * Samples for ScopeConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerScopeConnectionList.json
     */
    /**
     * Sample code: List Network Manager Scope Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listNetworkManagerScopeConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getScopeConnections().list("rg1", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
