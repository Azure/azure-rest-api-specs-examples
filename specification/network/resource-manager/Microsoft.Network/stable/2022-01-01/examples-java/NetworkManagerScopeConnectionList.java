import com.azure.core.util.Context;

/** Samples for ScopeConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerScopeConnectionList.json
     */
    /**
     * Sample code: List Network Manager Scope Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkManagerScopeConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getScopeConnections()
            .list("rg1", "testNetworkManager", null, null, Context.NONE);
    }
}
