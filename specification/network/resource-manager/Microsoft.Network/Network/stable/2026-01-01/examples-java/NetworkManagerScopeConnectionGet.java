
/**
 * Samples for ScopeConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkManagerScopeConnectionGet.json
     */
    /**
     * Sample code: Get Network Manager Scope Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getNetworkManagerScopeConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getScopeConnections().getWithResponse("rg1", "testNetworkManager",
            "TestScopeConnection", com.azure.core.util.Context.NONE);
    }
}
