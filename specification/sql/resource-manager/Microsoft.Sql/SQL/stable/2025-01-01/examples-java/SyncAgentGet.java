
/**
 * Samples for SyncAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncAgentGet.json
     */
    /**
     * Sample code: Get a sync agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncAgents().getWithResponse("syncagentcrud-65440", "syncagentcrud-8475",
            "syncagentcrud-3187", com.azure.core.util.Context.NONE);
    }
}
