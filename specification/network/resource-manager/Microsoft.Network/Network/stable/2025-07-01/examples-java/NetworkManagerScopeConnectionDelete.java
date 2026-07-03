
/**
 * Samples for ScopeConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerScopeConnectionDelete.json
     */
    /**
     * Sample code: Delete Network Manager Scope Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteNetworkManagerScopeConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getScopeConnections().deleteWithResponse("rg1", "testNetworkManager",
            "TestScopeConnection", com.azure.core.util.Context.NONE);
    }
}
