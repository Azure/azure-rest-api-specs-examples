
/**
 * Samples for SyncAgents ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncAgentListByServer.json
     */
    /**
     * Sample code: Get sync agents under a server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSyncAgentsUnderAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncAgents().listByServer("syncagentcrud-65440",
            "syncagentcrud-8475", com.azure.core.util.Context.NONE);
    }
}
