
/**
 * Samples for SyncAgents ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentListByServer.json
     */
    /**
     * Sample code: Get sync agents under a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getSyncAgentsUnderAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().listByServer("syncagentcrud-65440", "syncagentcrud-8475",
            com.azure.core.util.Context.NONE);
    }
}
