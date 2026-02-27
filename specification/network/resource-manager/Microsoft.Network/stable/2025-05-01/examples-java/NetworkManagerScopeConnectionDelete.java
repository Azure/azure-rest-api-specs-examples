
/**
 * Samples for ScopeConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerScopeConnectionDelete.json
     */
    /**
     * Sample code: Delete Network Manager Scope Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkManagerScopeConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getScopeConnections().deleteWithResponse("rg1", "testNetworkManager",
            "TestScopeConnection", com.azure.core.util.Context.NONE);
    }
}
