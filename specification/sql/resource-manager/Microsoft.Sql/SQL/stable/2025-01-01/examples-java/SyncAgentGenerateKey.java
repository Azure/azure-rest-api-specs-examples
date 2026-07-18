
/**
 * Samples for SyncAgents GenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentGenerateKey.json
     */
    /**
     * Sample code: Generate a sync agent key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void generateASyncAgentKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().generateKeyWithResponse("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
