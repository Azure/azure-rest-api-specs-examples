
/**
 * Samples for SyncAgents ListLinkedDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentGetLinkedDatabases.json
     */
    /**
     * Sample code: Get sync agent linked databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getSyncAgentLinkedDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().listLinkedDatabases("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
