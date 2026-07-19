
/**
 * Samples for SyncAgents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentDelete.json
     */
    /**
     * Sample code: Delete a sync agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteASyncAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().delete("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
